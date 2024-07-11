package main

import (
	"DumbiFadhil/edas-api/config"
	"DumbiFadhil/edas-api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Gin mode from the environment variable
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	// Initialize the Gin router
	router := config.SetupRouter()

	// Configure routes
	routes.SetupRoutes(router)

	// Print server start indicator
	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatal("Server failed to start:", err)
		}
	}()
	println("Server started on port 8080")

	// Block the main goroutine
	select {}
}
