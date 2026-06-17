package entity

import "time"

type Review struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"not null"`
	VehicleID int       `json:"vehicle_id" gorm:"not null"`
	RentalID  int       `json:"rental_id" gorm:"not null"`
	Rating    int16     `json:"rating" gorm:"not null"`
	Comment   *string   `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User    *User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Vehicle *Vehicle `json:"vehicle,omitempty" gorm:"foreignKey:VehicleID"`
	Rental  *Rental  `json:"rental,omitempty" gorm:"foreignKey:RentalID"`
}
