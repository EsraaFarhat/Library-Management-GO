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

		bookRoutes.POST("/", bookHandler.CreateBook)
		bookRoutes.GET("/", bookHandler.GetAllBooks)
		bookRoutes.GET("/:id", bookHandler.GetBook)
		bookRoutes.PUT("/:id", bookHandler.UpdateBook)
		bookRoutes.DELETE("/:id", bookHandler.DeleteBook)
	}
}
