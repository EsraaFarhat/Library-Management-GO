package error_handlers

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/utils/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleUserError handles errors specific to the UserHandler
func HandleUserError(c *gin.Context, err error) {
	var validationErr *handlers.ValidationError
	if errors.As(err, &validationErr) {
		handlers.RespondWithError(c, http.StatusBadRequest, validationErr)
		return
	}

	switch {
	case errors.Is(err, constants.ErrEmailTaken):
		handlers.RespondWithError(c, http.StatusConflict, err)
	case errors.Is(err, constants.ErrUserNotFound):
		handlers.RespondWithError(c, http.StatusNotFound, err)
	case errors.Is(err, constants.ErrInvalidInput):
		handlers.RespondWithError(c, http.StatusBadRequest, err)
	default:
		handlers.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
	}
}
