package dto

// ============================================================
// Request
// ============================================================

type UpdateProfileRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=255"`
	Phone string `json:"phone" validate:"required,min=10,max=20"`
}

type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required,min=6,max=255"`
	NewPassword     string `json:"new_password" validate:"required,min=6,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,max=255"`
}

// ============================================================
// Response
// ============================================================

type ProfileResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
