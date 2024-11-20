package main

import (
	"log"
	"subscription-tracker/internal/config"
	"subscription-tracker/internal/database"
	"subscription-tracker/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file in development
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db := database.InitDB(cfg)

	// Create and start server
	srv := server.New(db, cfg)

	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := srv.Start(":" + cfg.Server.Port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
