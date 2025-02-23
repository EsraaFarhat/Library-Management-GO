package services

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/repository"
	"library-management/internal/utils/auth"
	"library-management/internal/utils/mappers"
	"strings"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

// Create User (with hashed password)
func (s *AuthService) Register(req dto.UserRegisterRequest) (string, dto.UserResponse, error) {

	user := mappers.MapRegisterRequestToUser(req)
	// Convert email to lowercase
	user.Email = strings.ToLower(user.Email)

	// Check if the email already exists
	existingUser, _ := s.Repo.GetByEmail(user.Email, []string{"id"})
	if existingUser != nil {
		return "", dto.UserResponse{}, constants.ErrEmailTaken
	}

	// Hash password before saving
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return "", dto.UserResponse{}, err
	}
	user.Password = hashedPassword

	// Save to DB
	user, err = s.Repo.Create(user)
	if err != nil {
		return "", dto.UserResponse{}, err
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", dto.UserResponse{}, err
	}
	// Map user to response DTO
	userResponse := mappers.MapUserToResponse(user)
	return token, userResponse, nil
}

// Login (returns user if successful)
func (s *AuthService) Login(req dto.UserLoginRequest) (string, dto.UserResponse, error) {
	user := mappers.MapLoginRequestToUser(req)

	user, err := s.Repo.GetByEmail(user.Email, []string{"id", "name", "email", "password", "role", "created_at"})
	if err != nil {
		return "", dto.UserResponse{}, constants.ErrInvalidCredentials
	}

	// Verify password
	if !auth.CheckPasswordHash(req.Password, user.Password) {
		return "", dto.UserResponse{}, constants.ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", dto.UserResponse{}, err
	}
	// Map user to response DTO
	userResponse := mappers.MapUserToResponse(user)
	return token, userResponse, nil
}
