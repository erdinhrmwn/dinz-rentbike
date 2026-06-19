package entity

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Phone     string    `json:"phone" gorm:"not null"`
	Password  string    `json:"-" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null;default:customer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Reviews  []Review  `json:"reviews,omitempty" gorm:"foreignKey:UserID"`
	Rentals  []Rental  `json:"rentals,omitempty" gorm:"foreignKey:UserID"`
	Payments []Payment `json:"payments,omitempty" gorm:"foreignKey:UserID"`
}
