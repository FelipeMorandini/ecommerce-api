package main

import (
	"ecommerce-api/cmd/api"
	"ecommerce-api/internal/config"
	"ecommerce-api/internal/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it. Using environment variables.")
	}

	// Load configuration
	cfg := config.Load()

	// Connect to the database
	db, err := database.NewGormConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run database migrations
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	// Get the SQL DB instance for the API server
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB instance: %v", err)
	}
	defer sqlDB.Close()

	// Create and run the API server
	server := api.NewAPIServer(cfg.Server.Addr, sqlDB)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
