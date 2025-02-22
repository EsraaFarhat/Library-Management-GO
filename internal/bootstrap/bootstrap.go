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

	// Register routes
	routes.SetupUserRoutes(r, userHandler)

	return r
}
