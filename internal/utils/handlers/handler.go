package handlers

import (
	"encoding/json"
	"library-management/internal/constants"
	"library-management/internal/utils/validation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BindAndValidate binds the request body to a struct and validates it
func BindAndValidate(c *gin.Context, req interface{}) error {
	// Create a new decoder that disallows unknown fields
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields() // Prevent extra fields in JSON

	if err := decoder.Decode(req); err != nil {
		return constants.ErrInvalidInput
	}

	validate := validator.New()
	validate.RegisterValidation("password", validation.ValidatePassword)

	if err := validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return FormatValidationErrors(validationErrors, req)
		}
		return err
	}

	return nil
}
