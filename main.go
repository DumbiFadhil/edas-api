package main

import (
	"DumbiFadhil/edas-api/config"
	"DumbiFadhil/edas-api/routes"
	"DumbiFadhil/edas-api/services"
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

	// Initialize MongoDB connection
	mongoURI := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_DB_NAME")
	if mongoURI == "" || dbName == "" {
		log.Fatal("MONGODB_URI and MONGODB_DB_NAME must be set in the .env file")
	}
	services.ConnectToDB(mongoURI, dbName)

	// Test MongoDB connection
	err = services.TestDBConnection()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	} else {
		log.Println("Successfully connected to MongoDB!")
	}

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
