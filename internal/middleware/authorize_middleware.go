package middlewares

import (
	"net/http"

	"library-management/internal/constants"
	"library-management/internal/utils"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists || userRole != role {
			utils.RespondWithError(c, http.StatusForbidden, constants.ErrUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
