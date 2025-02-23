package dto

import "time"

// UserCreateRequest represents the input for user creation.
type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"omitempty,oneof=admin member"`
}

// UserUpdateRequest represents the input for user update.
type UserUpdateRequest struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty" validate:"email"`
	Password *string `json:"password,omitempty"`
	Role     *string `json:"role,omitempty" validate:"oneof=admin member"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
