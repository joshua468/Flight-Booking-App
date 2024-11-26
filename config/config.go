package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	ServerPort    string
	DatabaseURL   string
	JWTSecret     string
	PaymentAPIKey string
}

// LoadConfig loads the environment variables from the .env file in the root folder
func LoadConfig() (*Config, error) {
	// Build the path to the .env file in the root folder
	envPath := filepath.Join("..", ".env") // This points to the root folder
	absPath, _ := filepath.Abs(envPath)
	fmt.Println("Loading .env file from:", absPath)

	// Load the .env file into the environment
	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// Populate and return the configuration
	return &Config{
		ServerPort:    os.Getenv("SERVER_PORT"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		PaymentAPIKey: os.Getenv("PAYMENT_API_KEY"),
	}, nil
}
