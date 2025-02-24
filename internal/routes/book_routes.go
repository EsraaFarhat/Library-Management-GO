package routes

import (
	"library-management/internal/handlers"
	middlewares "library-management/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(r *gin.Engine, bookHandler *handlers.BookHandler) {
	bookRoutes := r.Group("/books")
	{
		bookRoutes.Use(middlewares.AuthMiddleware())

		bookRoutes.GET("/:id", bookHandler.GetBook)
		bookRoutes.GET("/", bookHandler.GetAllBooks)

		bookRoutes.Use(middlewares.RoleMiddleware("admin"))
		bookRoutes.POST("/", bookHandler.CreateBook)
		bookRoutes.PUT("/:id", bookHandler.UpdateBook)
		bookRoutes.DELETE("/:id", bookHandler.DeleteBook)
	}
}
