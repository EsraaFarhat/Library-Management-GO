package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"library-management/internal/bootstrap"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Warning: No .env file found, using default values.")
	}

	// Get port from environment variable (default to 8080 if not set)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := bootstrap.SetupServer()

	fmt.Printf("🚀 Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
