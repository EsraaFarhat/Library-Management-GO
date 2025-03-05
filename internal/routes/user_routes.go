package routes

import (
	"library-management/internal/constants"
	"library-management/internal/handlers"
	middlewares "library-management/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	userRoutes := r.Group("/users")
	{
		userRoutes.Use(middlewares.AuthMiddleware())
		userRoutes.Use(middlewares.RoleMiddleware(string(constants.Admin)))

		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
