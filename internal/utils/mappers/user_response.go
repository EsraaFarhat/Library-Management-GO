package mappers

import (
	"library-management/internal/dto"
	"library-management/internal/models"
)

// MapUserToResponse maps a models.User to a UserResponse
func MapUserToResponse(user *models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
}
