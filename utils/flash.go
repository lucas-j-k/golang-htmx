package utils

import (
	"encoding/base64"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SetFlash adds a flash message cookie, with a specified message which
// will be passed to the next template to be rendered
func SetFlash(ctx *fiber.Ctx, message string) {

	byteStringMessage := []byte(message)

	cookie := GenerateFlashCookie(byteStringMessage, time.Now().Add(1*time.Minute))

	ctx.Cookie(cookie)
}

// ClearFlash deletes the current flash message from the cookies
func ClearFlash(ctx *fiber.Ctx) {
	cookie := GenerateFlashCookie([]byte(""), time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	ctx.Cookie(cookie)

}

// WithFlash
// checks for a flash message on the current incoming request.
// If it finds one, it adds it to the req locals so we cn always pass to the template
// renderers. If a flash is found, it deletes the cookie so it does not get rendered
// more than once
func WithFlash() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		activeFlash := ctx.Cookies("flash-message")

		if activeFlash != "" {
			decoded, err := base64.URLEncoding.DecodeString(activeFlash)
			if err == nil {
				ctx.Locals("Flash", string(decoded))
			}
		}

		// Clear the flash cookie to avoid repeated requests
		ClearFlash(ctx)

		return ctx.Next()
	}
}

// GenerateFlashCookie creates a new flash cookie
func GenerateFlashCookie(message []byte, expires time.Time) *fiber.Cookie {

	cookie := new(fiber.Cookie)
	cookie.Name = "flash-message"
	cookie.Value = base64.URLEncoding.EncodeToString(message)
	cookie.Path = "/"
	cookie.HTTPOnly = true
	cookie.Expires = expires

	return cookie
}
