package link

import (
	sqlcservice "example.com/go-links-htmx/database"
	"github.com/go-playground/validator"
)

type LinkController struct {
	svc *sqlcservice.Queries
}

func New(svc *sqlcservice.Queries) LinkController {
	c := LinkController{
		svc: svc,
	}
	return c
}

type Link struct {
	URL        string `form:"url" validate:"required,url"`
	LinkTypeID int32  `form:"link_type_id"  validate:"required"`
	Published  bool   `form:"published"`
}

func (body *Link) Validate() map[string]bool {
	validate := validator.New()
	err := validate.Struct(body)
	if err == nil {
		return nil
	}

	valMap := make(map[string]bool)

	// Cast the err to ValidationErrors so we can get the field names that failed
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		valMap[field] = true
	}

	return valMap
}
