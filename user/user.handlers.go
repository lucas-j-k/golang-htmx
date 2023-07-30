package user

import (
	"example.com/go-links-htmx/auth"
	sqlcservice "example.com/go-links-htmx/database"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

// RegisterPostHandler receives and handles a POST request to create a user
// Redirects to login form, or renders the register template with errors
func (c *UserController) RegisterPostHandler(ctx *fiber.Ctx) error {

	// Try to parse the incoming req body
	var reqBody RegisterRequestBody

	err := ctx.BodyParser(&reqBody)
	if err != nil {
		return err
	}

	// Validate the incoming request
	valMap := reqBody.Validate()

	if len(valMap) > 0 {
		return ctx.Render("signup", fiber.Map{
			"Title":            "Create an account",
			"FormValues":       reqBody,
			"ValidationErrors": valMap,
		})
	}

	// Hash user password
	hashed, err := auth.HashPassword(reqBody.Password)
	if err != nil {
		return err
	}

	// Save new user into DB
	_, err = c.svc.InsertUser(ctx.Context(), sqlcservice.InsertUserParams{
		Username:     reqBody.Username,
		FirstName:    reqBody.FirstName,
		LastName:     reqBody.LastName,
		Email:        reqBody.Email,
		PasswordHash: hashed,
	})

	// Return redirect to login page
	ctx.Append("HX-Redirect", "/login")
	return ctx.Redirect("/login", fiber.StatusOK)

}

// LoginPostHandler receives and handles the post request from the login
// form. Redirect to the homepage if login is successful. Otherwise, return the
// loginform template with error details
func (c *UserController) LoginPostHandler(ctx *fiber.Ctx) error {

	// Unmarshall and validate request form body
	var reqBody LoginRequestBody

	err := ctx.BodyParser(&reqBody)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(reqBody)
	if err != nil {
		return ctx.Render("login", fiber.Map{
			"LoginFailed": true,
		})
	}

	userRows, err := c.svc.GetUserByEmail(ctx.Context(), reqBody.Email)

	if err != nil {
		return err
	}

	// User doesn't exist
	if len(userRows) < 1 {
		return ctx.Render("login", fiber.Map{
			"LoginFailed": true,
		})
	}

	user := userRows[0]

	// Verify the password hash
	validPassword := auth.PasswordsMatch(user.PasswordHash, reqBody.Password)
	if !validPassword {
		return ctx.Render("login", fiber.Map{
			"LoginFailed": true,
		})
	}

	// Generate a new session ID
	sessionId, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// Generate new session contents
	session := auth.UserSession{
		UserID:    user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// Persist the session to cache
	err = c.sess.SetSession(ctx.Context(), sessionId.String(), session)
	if err != nil {
		return err
	}

	// Set the session ID in the server-side cookie
	auth.SetAuthCookie(ctx, sessionId.String())

	// Return redirect to homepage
	ctx.Append("HX-Redirect", "/")
	return ctx.Redirect("/", fiber.StatusOK)

}

// LogoutUser clears the session cookie for the current user,
// and redirects to the homepage
func (c *UserController) LogoutPostHandler(ctx *fiber.Ctx) error {

	// If there is a session cookie, delete it
	sessionId := auth.GetSessionIdFromCookie(ctx)
	_, err := c.sess.GetSession(ctx.Context(), sessionId)

	if err != nil {
		return ctx.Redirect("/", fiber.StatusFound)
	}

	auth.ClearAuthCookie(ctx)

	// Return redirect to homepage
	return ctx.Redirect("/", fiber.StatusFound)
}
