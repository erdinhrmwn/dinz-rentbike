package dto

// ============================================================
// Request
// ============================================================

type CreateVehicleRequest struct {
	Type         string  `json:"type" validate:"required,oneof=motor mobil"`
	Brand        string  `json:"brand" validate:"required,max=255"`
	Name         string  `json:"name" validate:"required,max=255"`
	Category     string  `json:"category" validate:"required,max=100"`
	Description  *string `json:"description"`
	ImageURL     *string `json:"image_url"`
	PricePerHour float64 `json:"price_per_hour" validate:"required,gt=0"`
}

type UpdateVehicleRequest struct {
	Type         string  `json:"type" validate:"required,oneof=motor mobil"`
	Brand        string  `json:"brand" validate:"required,max=255"`
	Name         string  `json:"name" validate:"required,max=255"`
	Category     string  `json:"category" validate:"required,max=100"`
	Description  *string `json:"description"`
	ImageURL     *string `json:"image_url"`
	PricePerHour float64 `json:"price_per_hour" validate:"required,gt=0"`
	Status       string  `json:"status" validate:"required,oneof=available rented maintenance"`
}

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
