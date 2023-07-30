package link

import (
	"strconv"

	"example.com/go-links-htmx/auth"
	sqlcservice "example.com/go-links-htmx/database"
	"github.com/gofiber/fiber/v2"
)

// UpdateLinkPage
// Renders a form to update a single link
func (c *LinkController) UpdateLinkView(ctx *fiber.Ctx) error {

	// Get the target link ID from the url params
	idParam := ctx.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 32)

	// get user session from request locals
	user := ctx.Locals("User").(*auth.UserSession)

	if err != nil {
		return err
	}

	// get available link type options
	linkTypes, err := c.svc.ListLinkTypes(ctx.Context())
	if err != nil {
		return err
	}

	// get existing link details from DB
	link, err := c.svc.FindLinkById(ctx.Context(), sqlcservice.FindLinkByIdParams{
		ID:     int32(id),
		UserID: user.UserID,
	})

	if err != nil {
		return err
	}

	// render the insert form with the link types mapped in as options
	return ctx.Render("editLink", fiber.Map{
		"Title":     "Update Link",
		"Link":      link,
		"LinkTypes": linkTypes,
	}, "layouts/main")

}

// ManageLinksView
// Lists all the links for the logged in user
func (c *LinkController) ManageLinksView(ctx *fiber.Ctx) error {

	// get currently authed user from locals
	user := ctx.Locals("User").(*auth.UserSession)

	// get all links for user
	links, err := c.svc.ListLinksForUser(ctx.Context(), user.UserID)
	if err != nil {
		return err
	}

	// render the insert form with the link types mapped in as options
	return ctx.Render("listLinksAdmin", fiber.Map{
		"Title": "Manage Links",
		"Links": links,
	}, "layouts/main")

}

// CreateLinkPage
// Renders the create link form
func (c *LinkController) CreateLinkView(ctx *fiber.Ctx) error {

	// get available link type options
	linkTypes, err := c.svc.ListLinkTypes(ctx.Context())
	if err != nil {
		return err
	}

	// render the insert form with the link types mapped in as options
	return ctx.Render("insertLink", fiber.Map{
		"Title":     "Add a new link",
		"LinkTypes": linkTypes,
	}, "layouts/main")

}
