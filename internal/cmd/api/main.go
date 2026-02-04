package main

import (
	"database/sql"
	"errors"
	"fmt"
	httpDelivery "github.com/andriwhyu/simple-go-user-management/internal/delivery/http"
	"github.com/andriwhyu/simple-go-user-management/internal/infrastructure/database"
	"github.com/andriwhyu/simple-go-user-management/internal/repository"
	"github.com/andriwhyu/simple-go-user-management/internal/usecase"
	"github.com/andriwhyu/simple-go-user-management/internal/utils"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	// Initialize layers (Dependency Injection)
	// Repository layer
	userRepo := repository.NewUserRepository(db)

	// Usecase layer
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Delivery layer
	userHandler := httpDelivery.NewUserHandler(userUsecase)

	// Setup router
	router := httpDelivery.NewRouter(userHandler)

	// Server configuration
	port := utils.GetStringEnv("SERVER_PORT", "8080")
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("server starting on port %s", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("server shutting down...")
	if err := server.Close(); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server exited")
}
