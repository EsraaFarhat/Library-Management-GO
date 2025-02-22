package utils

import (
	"errors"

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

func FormatValidationErrors(errs validator.ValidationErrors) error {

	for _, e := range errs {
		// Create meaningful error messages
		switch e.Tag() {
		case "type":
			return &ValidationError{Message: e.Field() + " must be one of " + e.Param()} // Handle `oneof` tag
		case "required":
			return &ValidationError{Message: e.Field() + " is required"}
		case "email":
			return &ValidationError{Message: e.Field() + " must be a valid email"}
		case "min":
			return &ValidationError{Message: e.Field() + " must be at least " + e.Param() + " characters long"}
		case "oneof":
			return &ValidationError{Message: e.Field() + " must be one of " + e.Param()} // Handle `oneof` tag
		default:
			return &ValidationError{Message: e.Field() + " is invalid"}
		}
	}

	// Fallback error
	return errors.New("invalid input")
}
