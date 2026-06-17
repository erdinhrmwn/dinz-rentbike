package entity

import "time"

type Payment struct {
	ID               int        `json:"id" gorm:"primaryKey"`
	UserID           int        `json:"user_id" gorm:"not null"`
	RentalID         int        `json:"rental_id" gorm:"not null"`
	Amount           float64    `json:"amount" gorm:"not null"`
	Status           string     `json:"status" gorm:"not null;default:pending"`
	PaymentMethod    *string    `json:"payment_method"`
	XenditInvoiceID  *string    `json:"xendit_invoice_id" gorm:"unique"`
	XenditPaymentURL *string    `json:"xendit_payment_url"`
	PaidAt           *time.Time `json:"paid_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	User   *User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Rental *Rental `json:"rental,omitempty" gorm:"foreignKey:RentalID"`
}
