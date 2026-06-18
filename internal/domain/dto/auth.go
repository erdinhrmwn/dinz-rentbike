package dto

// ============================================================
// Request
// ============================================================

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Phone    string `json:"phone" validate:"required,min=10,max=20"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

// ============================================================
// Response
// ============================================================

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	User UserResponse
}
