package database

import (
	"database/sql"
	"fmt"
	"github.com/andriwhyu/simple-go-user-management/internal/utils"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed connecting to database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	return db, nil
}

func LoadConfigFromEnv() Config {
	return Config{
		Host:     utils.GetStringEnv("DB_HOST", "localhost"),
		Port:     utils.GetStringEnv("DB_PORT", "5432"),
		User:     utils.GetStringEnv("DB_USER", "admin"),
		Password: utils.GetStringEnv("DB_PASSWORD", "adminpassword"),
		DBName:   utils.GetStringEnv("DB_NAME", "user_management"),
		SSLMode:  utils.GetStringEnv("DB_SSLMODE", "disable"),
	}
}
