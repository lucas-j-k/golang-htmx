package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// ErrorView renders an error message in the UI. If any handlers return an error, this
// route handler will be called and rendered to the browser
func ErrorView(ctx *fiber.Ctx, err error) error {

	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	message := "An error occurred"

	if code == fiber.StatusNotFound {
		message = "Page not found"
	}

	return ctx.Render("error", fiber.Map{
		"Message": message,
	}, "layouts/main")

}
