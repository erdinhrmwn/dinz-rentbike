package dto

// ============================================================
// Response
// ============================================================

type VehicleResponse struct {
	ID           int              `json:"id"`
	Type         string           `json:"type"`
	Brand        string           `json:"brand"`
	Name         string           `json:"name"`
	Category     string           `json:"category"`
	Description  *string          `json:"description"`
	ImageURL     *string          `json:"image_url"`
	PricePerHour float64          `json:"price_per_hour"`
	Status       string           `json:"status"`
	CreatedAt    string           `json:"created_at"`
	UpdatedAt    string           `json:"updated_at"`
	Reviews      []ReviewResponse `json:"reviews,omitempty"`
}
