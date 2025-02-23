package handlers

import (
	"library-management/internal/constants"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BindAndValidate binds the request body to a struct and validates it
func BindAndValidate(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		return constants.ErrInvalidInput
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return FormatValidationErrors(validationErrors, req)
		}
		return err
	}

	return nil
}
