package user

import (
	"example.com/go-links-htmx/auth"
	"github.com/gofiber/fiber/v2"
)

// RegisterView renders the user registration form
func (c *UserController) RegisterView(ctx *fiber.Ctx) error {
	return ctx.Render("signup", fiber.Map{
		"Title": "Create an account",
	}, "layouts/main")
}

// RegisterView renders the user login form
func (c *UserController) LoginView(ctx *fiber.Ctx) error {

	// If user is already logged in, redirect to home route
	user := ctx.Locals("User")

	if user != nil {
		return ctx.Redirect("/", fiber.StatusFound)
	}

	// Render the html template and pass it the data
	return ctx.Render("login", fiber.Map{
		"Title": "Login",
	}, "layouts/main")
}

// UserProfileView renders the user settings / information page
func (c *UserController) UserProfileView(ctx *fiber.Ctx) error {

	// get user session from request locals
	user := ctx.Locals("User").(*auth.UserSession)

	// get user profile details from the DB
	userRecord, err := c.svc.GetUserProfile(ctx.Context(), user.UserID)

	if err != nil {
		return err
	}

	// render the insert form with the link types mapped in as options
	return ctx.Render("profile", fiber.Map{
		"Title": "User Profile",
		"User":  userRecord,
	}, "layouts/main")
}

// UserPublicProfileView
// Renders the public link list for the user
func (c *UserController) UserPublicProfileView(ctx *fiber.Ctx) error {

	// get user id from path
	username := ctx.Params("username")

	// fetch user details for the user
	userRows, err := c.svc.GetUserByUsername(ctx.Context(), username)

	if len(userRows) < 1 {
		return err
	}

	user := userRows[0]

	// fetch links for this user
	links, err := c.svc.ListPublicLinksForUser(ctx.Context(), user.ID)

	if err != nil {
		return err
	}

	// render the insert form with the link types mapped in as options
	return ctx.Render("publicProfile", fiber.Map{
		"Username": user.Username,
		"Links":    links,
	}, "layouts/public")
}
