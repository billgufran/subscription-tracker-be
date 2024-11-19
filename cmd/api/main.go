package main

import (
	"log"
	"os"

	"subscription-tracker/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db := database.InitDB()

	// Create server
	server := NewServer(db)

	// Get port from environment variable or use default
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := server.Start(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
