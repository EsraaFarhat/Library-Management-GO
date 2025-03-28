package routes

import (
	"library-management/internal/constants"
	"library-management/internal/handlers"
	middlewares "library-management/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupBorrowRoutes(r *gin.Engine, borrowHandler *handlers.BorrowHandler) {
	borrowRoutes := r.Group("/borrows")
	{
		borrowRoutes.Use(middlewares.AuthMiddleware())
		borrowRoutes.POST("/", borrowHandler.BorrowBook)
		borrowRoutes.PATCH("/return", borrowHandler.ReturnBook)
		// Get borrowed books for the logged-in user
		borrowRoutes.GET("/", borrowHandler.GetMyBorrows)

		borrowRoutes.Use(middlewares.RoleMiddleware(string(constants.Admin)))
		// borrowRoutes.GET("/", borrowHandler.GetBorrowRecords)
		borrowRoutes.GET("/users/:user_id", borrowHandler.GetUserBorrows)
	}
}
