package services

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/models"
	"library-management/internal/repository"
	"library-management/internal/utils"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Create User (with hashed password)
func (s *UserService) CreateUser(req dto.UserCreateRequest) (*models.User, error) {
	user := utils.MapCreateRequestToUser(req)

	// Validate struct
	if err := validate.Struct(req); err != nil {
		// Extract validation errors and return the first error
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return nil, utils.FormatValidationErrors(validationErrors)
		}
		return nil, err
	}

	// Convert email to lowercase
	user.Email = strings.ToLower(user.Email)

	// Check if the email already exists
	existingUser, _ := s.Repo.GetByEmail(user.Email)
	if existingUser != nil {
		return nil, constants.ErrEmailTaken
	}

	// Hash password before saving
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	user, err = s.Repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Get User by ID
func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.Repo.GetByID(id)
}

// Update User
func (s *UserService) UpdateUser(id uint, req dto.UserUpdateRequest) (*models.User, error) {
	user, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, constants.ErrUserNotFound
	}

	// Update user fields
	utils.UpdateUserFromDTO(user, req)

	// Validate struct
	if err := validate.Struct(req); err != nil {
		// Extract validation errors and return the first error
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return nil, utils.FormatValidationErrors(validationErrors)
		}
		return nil, err
	}

	// Update only fields that are provided
	// Convert email to lowercase
	if user.Email != "" {
		user.Email = strings.ToLower(user.Email)

		// Check if the email is already in use by another user
		existingUser, _ := s.Repo.GetByEmail(user.Email)
		if existingUser != nil && existingUser.ID != id {
			return nil, constants.ErrEmailTaken
		}

	}
	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	err = s.Repo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Delete User
func (s *UserService) DeleteUser(id uint) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return constants.ErrUserNotFound
	}

	return s.Repo.Delete(id)
}
