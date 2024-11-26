package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/flight-booking-app/config"
	"github.com/joshua468/flight-booking-app/handlers"
	"github.com/joshua468/flight-booking-app/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Setting the client_encoding to UTF-8
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get raw DB connection: %v", err)
	}
	_, err = sqlDB.Exec("SET client_encoding TO 'UTF8'")
	if err != nil {
		log.Fatalf("Failed to set client encoding: %v", err)
	}

	log.Println("Database connection established with UTF-8 encoding")
	return db, nil
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	dbConn, err := ConnectDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		sqlDB, _ := dbConn.DB()
		sqlDB.Close()
	}()

	// Initialize services
	flightService := services.NewFlightService(dbConn)

	// Initialize Gin router
	r := gin.Default()

	// Handlers
	authHandler := handlers.NewAuthHandler(dbConn, cfg.JWTSecret)
	flightHandler := handlers.NewFlightHandler(flightService)

	// Register routes
	r.POST("/signup", authHandler.Signup)
	r.POST("/login", authHandler.Login)
	r.GET("/flights", flightHandler.SearchFlights)
	r.POST("/book-flight", flightHandler.BookFlight)

	// Start server
	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
