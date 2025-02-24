package handlers

import (
	"library-management/internal/constants"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Standard API response
func RespondWithError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{"error": err.Error()})
}

// Standard success response
func RespondWithSuccess(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{"data": data})
}

// ValidationError represents a structured validation error
type ValidationError struct {
	Message string
}

// Implement the error interface
func (v *ValidationError) Error() string {
	return v.Message
}

func FormatValidationErrors(errs validator.ValidationErrors, structInstance interface{}) error {
	reflected := reflect.TypeOf(structInstance)

	for _, e := range errs {
		// Get the JSON tag
		field, found := reflected.Elem().FieldByName(e.StructField())
		jsonTag := e.Field() // Default to Go field name if no tag found
		if found {
			jsonTag = field.Tag.Get("json")
			if jsonTag != "" {
				jsonTag = strings.Split(jsonTag, ",")[0]
			} else {
				jsonTag = e.Field()
			}
		}
		// Create meaningful error messages
		switch e.Tag() {
		case "type":
			return &ValidationError{Message: jsonTag + " must be one of " + e.Param()} // Handle `oneof` tag
		case "required":
			return &ValidationError{Message: jsonTag + " is required"}
		case "email":
			return &ValidationError{Message: jsonTag + " must be a valid email"}
		case "min":
			return &ValidationError{Message: jsonTag + " must be at least " + e.Param() + " characters long"}
		case "oneof":
			return &ValidationError{Message: jsonTag + " must be one of " + e.Param()} // Handle `oneof` tag
		case "password":
			return &ValidationError{Message: jsonTag + " must be at least 8 characters long, contain 1 uppercase, 1 lowercase, 1 number, and 1 special character"}
		default:
			return &ValidationError{Message: jsonTag + " is invalid"}
		}
	}

	// Fallback error
	return constants.ErrInvalidInput
}
