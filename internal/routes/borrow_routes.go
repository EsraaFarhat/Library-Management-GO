package routes

import (
	"library-management/internal/handlers"
	middlewares "library-management/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupBorrowRoutes(r *gin.Engine, borrowHandler *handlers.BorrowHandler) {
	borrowRoutes := r.Group("/borrows")
	{
		borrowRoutes.Use(middlewares.AuthMiddleware())

		borrowRoutes.POST("/", borrowHandler.BorrowBook)
		borrowRoutes.POST("/return", borrowHandler.ReturnBook)
		borrowRoutes.GET("/user/:user_id", borrowHandler.GetUserBorrows)
	}
}