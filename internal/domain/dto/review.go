package dto

// ============================================================
// Request
// ============================================================

type CreateReviewRequest struct {
	RentalID int    `json:"rental_id" validate:"required"`
	Rating   int    `json:"rating" validate:"required,min=1,max=5"`
	Comment  string `json:"comment" validate:"max=500"`
}

// ============================================================
// Response
// ============================================================

type ReviewResponse struct {
	ID        int              `json:"id"`
	UserID    int              `json:"user_id"`
	VehicleID int              `json:"vehicle_id"`
	RentalID  int              `json:"rental_id"`
	Rating    int              `json:"rating"`
	Comment   *string          `json:"comment"`
	CreatedAt string           `json:"created_at"`
	UpdatedAt string           `json:"updated_at"`
	User      *UserResponse    `json:"user,omitempty"`
	Vehicle   *VehicleResponse `json:"vehicle,omitempty"`
}
