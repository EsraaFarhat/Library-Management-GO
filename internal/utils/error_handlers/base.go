package error_handlers

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/utils/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBaseError(c *gin.Context, err error) {
	var validationErr *handlers.ValidationError
	if errors.As(err, &validationErr) {
		handlers.RespondWithError(c, http.StatusBadRequest, validationErr)
		return
	}

	// Default to internal server error
	handlers.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
}
