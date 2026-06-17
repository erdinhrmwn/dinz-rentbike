package entity

import "time"

type Vehicle struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Type         string    `json:"type" gorm:"not null"`
	Brand        string    `json:"brand" gorm:"not null"`
	Name         string    `json:"name" gorm:"not null"`
	Category     string    `json:"category" gorm:"not null"`
	Description  *string   `json:"description"`
	ImageURL     *string   `json:"image_url"`
	PricePerHour float64   `json:"price_per_hour" gorm:"not null"`
	Status       string    `json:"status" gorm:"not null;default:available"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
