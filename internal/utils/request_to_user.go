package utils

import (
	"library-management/internal/dto"
	"library-management/internal/models"
)

func ToUser(req dto.UserCreateRequest) *models.User {
	return &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}

func UpdateUserFromDTO(user *models.User, req dto.UserUpdateRequest) {
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.Role != "" {
		user.Role = req.Role
	}
}
