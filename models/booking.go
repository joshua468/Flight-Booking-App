package models

import "time"

type Booking struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	FlightID  uint   `gorm:"not null"`
	Status    string `gorm:"default:confirmed"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
