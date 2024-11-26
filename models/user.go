package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FullName  string `gorm:"not null"`
	Role      string `gorm:"default:customer"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
