package error_handlers

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/utils/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleBorrowError handles errors specific to the BorrowHandler
func HandleBorrowError(c *gin.Context, err error) {
	var validationErr *handlers.ValidationError
	if errors.As(err, &validationErr) {
		handlers.RespondWithError(c, http.StatusBadRequest, validationErr)
		return
	}

	switch {
	case errors.Is(err, constants.ErrUserNotFound),
		errors.Is(err, constants.ErrBookNotFound),
		errors.Is(err, constants.ErrBorrowNotFound):
		handlers.RespondWithError(c, http.StatusBadRequest, err)
	default:
		handlers.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
	}
}
