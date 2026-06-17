package entity

import "time"

type Rental struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	UserID     int       `json:"user_id" gorm:"not null"`
	VehicleID  int       `json:"vehicle_id" gorm:"not null"`
	StartTime  time.Time `json:"start_time" gorm:"not null"`
	EndTime    time.Time `json:"end_time" gorm:"not null"`
	TotalHours int       `json:"total_hours" gorm:"not null"`
	TotalPrice float64   `json:"total_price" gorm:"not null"`
	Status     string    `json:"status" gorm:"not null;default:pending"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	User    *User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Vehicle *Vehicle `json:"vehicle,omitempty" gorm:"foreignKey:VehicleID"`

	Review  []Review  `json:"reviews,omitempty" gorm:"foreignKey:RentalID"`
	Payment []Payment `json:"payments,omitempty" gorm:"foreignKey:RentalID"`
}
