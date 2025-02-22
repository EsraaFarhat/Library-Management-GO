package dto

// UserCreateRequest represents the input for user creation.
type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"omitempty,oneof=admin member"`
}

// UserUpdateRequest represents the input for user update.
type UserUpdateRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty" validate:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty" validate:"oneof=admin member"`
}
