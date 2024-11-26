package services

import (
	"errors"

	"github.com/joshua468/flight-booking-app/models"
	"gorm.io/gorm"
)

type FlightService struct {
	db *gorm.DB
}

func NewFlightService(db *gorm.DB) *FlightService {
	return &FlightService{db: db}
}

// SearchFlights searches for flights based on origin and destination
func (s *FlightService) SearchFlights(origin, destination string) ([]models.Flight, error) {
	var flights []models.Flight
	// Search flights by origin and destination
	if err := s.db.Where("origin = ? AND destination = ?", origin, destination).Find(&flights).Error; err != nil {
		return nil, err // Return any database-related errors
	}
	return flights, nil
}

// BookFlight books a flight for the given user
func (s *FlightService) BookFlight(flightID, userID uint) (*models.Booking, error) {
	var flight models.Flight
	// Retrieve the flight by ID
	if err := s.db.First(&flight, flightID).Error; err != nil {
		// Handle case where flight is not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("flight not found")
		}
		// Handle other types of errors (e.g., database connection issues)
		return nil, errors.New("failed to fetch flight details")
	}

	// Check if there are available seats
	if flight.SeatsAvailable <= 0 {
		return nil, errors.New("no seats available")
	}

	// Decrease the seat count (ensure it doesn't go below 0)
	flight.SeatsAvailable -= 1
	if err := s.db.Save(&flight).Error; err != nil {
		// Handle any error that occurs while updating the flight
		return nil, errors.New("failed to update seat availability")
	}

	// Create a booking
	booking := models.Booking{
		UserID:   userID,
		FlightID: flightID,
	}
	if err := s.db.Create(&booking).Error; err != nil {
		// Handle error while creating the booking
		return nil, errors.New("failed to create booking")
	}

	// Return the booking details on success
	return &booking, nil
}
