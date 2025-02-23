package bootstrap

import (
	"library-management/config"
	"library-management/internal/handlers"
	"library-management/internal/repository"
	"library-management/internal/routes"
	"library-management/internal/services"

	"github.com/gin-gonic/gin"
)

// Initialize and return Gin router
func SetupServer() *gin.Engine {
	db := config.ConnectDatabase()
	r := gin.Default()

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	bookRepo := repository.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	borrowRepo := repository.NewBorrowRepository(db)
	borrowService := services.NewBorrowService(borrowRepo, bookRepo, userRepo)
	borrowHandler := handlers.NewBorrowHandler(borrowService)

	// Register routes
	routes.SetupUserRoutes(r, userHandler)
	routes.SetupAuthRoutes(r, authHandler)
	routes.SetupBookRoutes(r, bookHandler)
	routes.SetupBorrowRoutes(r, borrowHandler)

	return r
}
