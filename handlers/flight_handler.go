package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/flight-booking-app/services"
)

type FlightHandler struct {
	flightService *services.FlightService
}

func NewFlightHandler(service *services.FlightService) *FlightHandler {
	return &FlightHandler{flightService: service}
}

func (h *FlightHandler) SearchFlights(c *gin.Context) {
	origin := c.Query("origin")
	destination := c.Query("destination")

	// Input validation
	if origin == "" || destination == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Origin and destination are required"})
		return
	}

	flights, err := h.flightService.SearchFlights(origin, destination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flights"})
		return
	}

	// If no flights found
	if len(flights) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No flights found"})
		return
	}

	c.JSON(http.StatusOK, flights)
}

func (h *FlightHandler) BookFlight(c *gin.Context) {
	var req struct {
		FlightID uint `json:"flight_id" binding:"required"`
		UserID   uint `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Attempt to book the flight
	booking, err := h.flightService.BookFlight(req.FlightID, req.UserID)
	if err != nil {
		// Detailed error handling based on the message
		if err.Error() == "flight not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		} else if err.Error() == "no seats available" {
			c.JSON(http.StatusConflict, gin.H{"error": "No seats available"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book flight"})
		}
		return
	}

	// Return the booking details on success
	c.JSON(http.StatusCreated, booking)
}
