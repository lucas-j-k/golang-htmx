package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// SetAuthCookie
// creates a cookie to save the current user session ID
func SetAuthCookie(ctx *fiber.Ctx, sessionId string) {
	cookie := new(fiber.Cookie)
	cookie.Name = "user-session"
	cookie.Value = sessionId
	cookie.Domain = "localhost"
	cookie.Path = "/"
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(24 * time.Hour)

	// Set cookie on fiber context
	ctx.Cookie(cookie)
}

// ClearAuthCookie
// Removes the current user session cookie
func ClearAuthCookie(ctx *fiber.Ctx) {
	ctx.ClearCookie("user-session")
}

// GetSessionIdFromCookie
// Retrieves the current session cookie from the request context
func GetSessionIdFromCookie(ctx *fiber.Ctx) string {
	value := ctx.Cookies("user-session")
	return value
}
