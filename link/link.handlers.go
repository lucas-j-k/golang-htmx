package link

import (
	"strconv"

	"example.com/go-links-htmx/auth"
	sqlcservice "example.com/go-links-htmx/database"
	"github.com/gofiber/fiber/v2"

	"example.com/go-links-htmx/utils"
)

// UpdateLinkHandler
// Handles a form request to update an existing link in the database
func (c *LinkController) UpdateLinkHandler(ctx *fiber.Ctx) error {

	// Try to parse the incoming req body
	var reqBody Link
	err := ctx.BodyParser(&reqBody)
	if err != nil {
		return err
	}

	// Validate the incoming request
	validationMap := reqBody.Validate()

	// Get the target link ID from the url params
	targetIdStr := ctx.Params("id")
	intId, err := strconv.ParseInt(targetIdStr, 10, 32)

	if err != nil {
		return err
	}

	int32Id := int32(intId)

	// if validation errors found, re-render the form with the errors
	if len(validationMap) > 0 {

		linkTypes, err := c.svc.ListLinkTypes(ctx.Context())
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}

		displayLink := sqlcservice.FindLinkByIdRow{
			ID:         int32Id,
			Url:        reqBody.URL,
			LinkTypeID: reqBody.LinkTypeID,
			Published:  reqBody.Published,
		}

		return ctx.Render("editLink", fiber.Map{
			"Title":            "Add a new link",
			"LinkTypes":        linkTypes,
			"Link":             displayLink,
			"ValidationErrors": validationMap,
		})
	}

	// Save new link into DB
	err = c.svc.UpdateLink(ctx.Context(), sqlcservice.UpdateLinkParams{
		Url:        reqBody.URL,
		LinkTypeID: reqBody.LinkTypeID,
		Published:  reqBody.Published,
		ID:         int32Id,
	})

	if err != nil {
		return err
	}

	// Set a flash message to render on redirected page to indicate successful edit
	utils.SetFlash(ctx, "Link updated successfully")

	// Return redirect to links list page
	ctx.Append("HX-Redirect", "/admin/links")
	return ctx.Redirect("/admin/links", fiber.StatusOK)
}

// DeleteLinkHandler deletes a link by ID
func (c *LinkController) DeleteLinkHandler(ctx *fiber.Ctx) error {

	targetIdStr := ctx.Params("id")
	intId, err := strconv.ParseInt(targetIdStr, 10, 32)

	if err != nil {
		return err
	}

	int32Id := int32(intId)

	delErr := c.svc.DeleteLink(ctx.Context(), int32Id)
	if delErr != nil {
		return err
	}

	// Set flash confirmation message
	utils.SetFlash(ctx, "Link deleted successfully")

	// redirect back to the links page
	return ctx.Redirect("/admin/links", fiber.StatusFound)
}

// CreateLinkHandler
// Handles a form request to insert a new link to the DB
func (c *LinkController) CreateLinkHandler(ctx *fiber.Ctx) error {

	// Try to parse the incoming req body
	var reqBody Link
	err := ctx.BodyParser(&reqBody)
	if err != nil {
		return err
	}

	// Validate the incoming request
	validationMap := reqBody.Validate()

	// if validation errors found, re-render the form with the errors
	if len(validationMap) > 0 {

		linkTypes, err := c.svc.ListLinkTypes(ctx.Context())
		if err != nil {
			return err
		}

		return ctx.Render("insertLink", fiber.Map{
			"Title":            "Add a new link",
			"LinkTypes":        linkTypes,
			"Link":             reqBody,
			"ValidationErrors": validationMap,
		})
	}

	// get user session from request locals
	user := ctx.Locals("User").(*auth.UserSession)

	// Save new link into DB
	_, err = c.svc.InsertLink(ctx.Context(), sqlcservice.InsertLinkParams{
		UserID:     user.UserID,
		Url:        reqBody.URL,
		LinkTypeID: reqBody.LinkTypeID,
		Published:  reqBody.Published,
	})

	// Set flash to notify of success
	utils.SetFlash(ctx, "Link added successfully")

	// Return redirect to links list page
	ctx.Append("HX-Redirect", "/admin/links")
	return ctx.Redirect("/admin/links", fiber.StatusOK)
}
