package middlewares

import (
	"log"
	"net/http"

	"library-management/internal/constants"
	"library-management/internal/utils"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		log.Println(c.Get("role"))
		if !exists || userRole != role {
			utils.RespondWithError(c, http.StatusForbidden, constants.ErrUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
