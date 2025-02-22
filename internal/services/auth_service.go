package services

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/repository"
	"library-management/internal/utils"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

// Create User (with hashed password)
func (s *AuthService) Register(req dto.UserRegisterRequest) (string, error) {
	// Validate struct
	if err := validate.Struct(req); err != nil {
		// Extract validation errors and return the first error
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return "", utils.FormatValidationErrors(validationErrors)
		}
		return "", err
	}

	user := utils.MapRegisterRequestToUser(req)
	// Convert email to lowercase
	user.Email = strings.ToLower(user.Email)

	// Check if the email already exists
	existingUser, _ := s.Repo.GetByEmail(user.Email)
	if existingUser != nil {
		return "", constants.ErrEmailTaken
	}

	// Hash password before saving
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	// Save to DB
	user, err = s.Repo.Create(user)
	if err != nil {
		return "", err
	}

	// Generate JWT token
	return utils.GenerateToken(user.ID, user.Role)
}

// Login (returns user if successful)
func (s *AuthService) Login(req dto.UserLoginRequest) (string, error) {
	user := utils.MapLoginRequestToUser(req)

	// Validate struct
	if err := validate.Struct(req); err != nil {
		// Extract validation errors and return the first error
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return "", utils.FormatValidationErrors(validationErrors)
		}
		return "", err
	}

	user, err := s.Repo.GetByEmail(user.Email)
	if err != nil {
		return "", constants.ErrInvalidCredentials
	}

	// Verify password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", constants.ErrInvalidCredentials
	}

	// Generate JWT token
	return utils.GenerateToken(user.ID, user.Role)
}
