package dto

// ============================================================
// Request
// ============================================================

type CreateRentalRequest struct {
	VehicleID int    `json:"vehicle_id" validate:"required"`
	StartTime string `json:"start_time" validate:"required"`
	EndTime   string `json:"end_time" validate:"required"`
}

type CancelRentalRequest struct {
	RentalID int `json:"rental_id" validate:"required"`
}

// ============================================================
// Response
// ============================================================

type RentalResponse struct {
	ID         int              `json:"id"`
	UserID     int              `json:"user_id"`
	VehicleID  int              `json:"vehicle_id"`
	StartTime  string           `json:"start_time"`
	EndTime    string           `json:"end_time"`
	TotalHours int              `json:"total_hours"`
	TotalPrice float64          `json:"total_price"`
	Status     string           `json:"status"`
	CreatedAt  string           `json:"created_at"`
	UpdatedAt  string           `json:"updated_at"`
	Vehicle    VehicleResponse  `json:"vehicle,omitempty"`
	Payment    *PaymentResponse `json:"payment,omitempty"`
	Review     *ReviewResponse  `json:"review,omitempty"`
}
