package models

type Flight struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	Origin         string `gorm:"not null"`
	Destination    string `gorm:"not null"`
	SeatsAvailable int    `gorm:"not null"`
	PricePerSeat   float64
}
