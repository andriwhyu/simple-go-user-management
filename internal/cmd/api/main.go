package main

import (
	"github.com/andriwhyu/simple-go-user-management/internal/infrastructure/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Database connection
	dbConfig := database.LoadConfigFromEnv()
	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize layers (Dependency Injection)
	// Repository layer
	//userRepo := repository.NewUserRepository()
}
