package mappers

import (
	"library-management/internal/dto"
	"library-management/internal/models"
)

func MapCreateRequestToUser(req dto.UserCreateRequest) *models.User {
	return &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}

func MapRegisterRequestToUser(req dto.UserRegisterRequest) *models.User {
	return &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func MapLoginRequestToUser(req dto.UserLoginRequest) *models.User {
	return &models.User{
		Email:    req.Email,
		Password: req.Password,
	}
}

func UpdateUserFromDTO(user *models.User, req dto.UserUpdateRequest) {
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Password != nil {
		user.Password = *req.Password
	}
	if req.Role != nil {
		user.Role = *req.Role
	}
}
