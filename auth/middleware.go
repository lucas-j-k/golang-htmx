package auth

import (
	"github.com/gofiber/fiber/v2"
)

// WithUser attempts to get an active user session, and add the current user to the Fiber context locals,
// to make the user object available in all upstream handlers
func WithUser(session SessionManager) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// Retrieve the user struct from Redis, using the sessionID in the request cookie
		sessionId := GetSessionIdFromCookie(ctx)
		user, err := session.GetSession(ctx.Context(), sessionId)

		if err != nil {
			// no valid session found, pass through
			return ctx.Next()
		}

		// we have a current session for user, set the user struct in the Locals of the current context
		ctx.Locals("User", user)

		// Pass request to the next handler in the chain
		return ctx.Next()
	}
}

// AuthGuard
// redirects the user to the login page if no valid user session was found
func AuthGuard(ctx *fiber.Ctx) error {

	htmxHeader := ctx.Get("HX-Request")

	user := ctx.Locals("User")

	// if request is coming from htmx, return a client side HTMX redirect
	if user == nil && htmxHeader == "true" {
		ctx.Append("HX-Redirect", "/login")
		return ctx.Redirect("/login", fiber.StatusOK)
	}

	// request is not triggered from HTMX - return full server side redirect
	if user == nil {
		return ctx.Redirect("/login", fiber.StatusFound)
	}

	return ctx.Next()
}
