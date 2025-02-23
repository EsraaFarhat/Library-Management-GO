package middlewares

import (
	"net/http"
	"strings"

	"library-management/internal/constants"
	"library-management/internal/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {

			utils.RespondWithError(c, http.StatusUnauthorized, constants.ErrMissingAuthHeader)
			c.Abort()
			return
		}

		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, constants.ErrInvalidTokenFormat)
			c.Abort()
			return
		}

		// Verify token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, constants.ErrInvalidOrExpiredToken)
			c.Abort()
			return
		}

		// Store user ID in context
		c.Set("user", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
