package dto

// ============================================================
// Request
// ============================================================

type CreatePaymentRequest struct {
	RentalID int `json:"rental_id" validate:"required"`
}

type CancelPaymentRequest struct {
	RentalID int `json:"rental_id" validate:"required"`
}

// ============================================================
// Response
// ============================================================

type PaymentResponse struct {
	ID               int     `json:"id"`
	UserID           int     `json:"user_id"`
	RentalID         int     `json:"rental_id"`
	Amount           float64 `json:"amount"`
	Status           string  `json:"status"`
	XenditInvoiceID  *string `json:"xendit_invoice_id"`
	XenditPaymentURL *string `json:"xendit_payment_url"`
	PaidAt           *string `json:"paid_at"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}
