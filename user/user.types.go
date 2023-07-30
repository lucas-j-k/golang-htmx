package user

import (
	"example.com/go-links-htmx/auth"
	sqlcservice "example.com/go-links-htmx/database"
	"github.com/go-playground/validator"
)

type UserController struct {
	svc  *sqlcservice.Queries
	sess auth.SessionManager
}

func New(svc *sqlcservice.Queries, sess auth.SessionManager) UserController {
	c := UserController{
		svc:  svc,
		sess: sess,
	}
	return c
}

// LoginRequestBody defines the incoming form data for a login request
type LoginRequestBody struct {
	Email    string `form:"email"  validate:"required,email"`
	Password string `form:"password"  validate:"required,min=8"`
}

// RegisterRequestBody defines the incoming form data for a register
// user request
type RegisterRequestBody struct {
	Username        string `form:"username" validate:"required,min=6,max=50"`
	FirstName       string `form:"firstName" validate:"required"`
	LastName        string `form:"lastName"  validate:"required"`
	Password        string `form:"password"  validate:"required,min=8"`
	PasswordConfirm string `form:"passwordConfirm"  validate:"required,min=8"`
	Email           string `form:"email"  validate:"required,email"`
}

func (body *RegisterRequestBody) Validate() map[string]bool {
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

	// password fields must match
	if body.Password != body.PasswordConfirm {
		valMap["PasswordConfirm"] = true
	}

	return valMap
}
