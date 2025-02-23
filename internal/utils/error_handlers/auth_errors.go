package error_handlers

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/utils/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAuthError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, constants.ErrEmailTaken):
		handlers.RespondWithError(c, http.StatusConflict, err)
	case errors.Is(err, constants.ErrInvalidCredentials):
		handlers.RespondWithError(c, http.StatusUnauthorized, err)
	default:
		handlers.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
	}
}
