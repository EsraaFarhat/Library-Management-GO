package error_handlers

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/utils/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBookError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, constants.ErrISBNExists):
		handlers.RespondWithError(c, http.StatusConflict, err)
	case errors.Is(err, constants.ErrBookNotFound):
		handlers.RespondWithError(c, http.StatusNotFound, err)
	default:
		handlers.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
	}
}
