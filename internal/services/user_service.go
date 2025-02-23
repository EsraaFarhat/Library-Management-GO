package services

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/repository"
	"library-management/internal/utils/auth"
	"library-management/internal/utils/mappers"
	"strings"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Create User (with hashed password)
func (s *UserService) CreateUser(req dto.UserCreateRequest) (dto.UserResponse, error) {
	user := mappers.MapCreateRequestToUser(req)

	// Convert email to lowercase
	user.Email = strings.ToLower(user.Email)

	// Check if the email already exists
	existingUser, _ := s.Repo.GetByEmail(user.Email, []string{"id"})
	if existingUser != nil {
		return dto.UserResponse{}, constants.ErrEmailTaken
	}

	// Hash password before saving
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return dto.UserResponse{}, err
	}
	user.Password = hashedPassword
	user, err = s.Repo.Create(user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	// Map user to response DTO
	userResponse := mappers.MapUserToResponse(user)
	return userResponse, nil
}

// Get User by ID
func (s *UserService) GetUser(id uint, fields []string) (dto.UserResponse, error) {
	user, err := s.Repo.GetByID(id, fields)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Map user to response DTO
	userResponse := mappers.MapUserToResponse(user)
	return userResponse, nil
}

// Get All Users
func (s *UserService) GetAllUsers(page, limit int, fields []string) ([]dto.UserResponse, int64, error) {
	// Fetch users from the repository
	users, total, err := s.Repo.GetAll(page, limit, fields)
	if err != nil {
		return nil, 0, err
	}

	// Map each models.User to dto.UserResponse
	userResponses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = mappers.MapUserToResponse(&user)
	}

	return userResponses, total, nil
}

// Update User
func (s *UserService) UpdateUser(id uint, req dto.UserUpdateRequest) (dto.UserResponse, error) {
	user, err := s.Repo.GetByID(id, []string{})
	if err != nil {
		return dto.UserResponse{}, constants.ErrUserNotFound
	}

	// Update user fields
	mappers.UpdateUserFromDTO(user, req)

	// Update only fields that are provided
	// Convert email to lowercase
	if user.Email != "" {
		user.Email = strings.ToLower(user.Email)

		// Check if the email is already in use by another user
		existingUser, _ := s.Repo.GetByEmail(user.Email, []string{"id"})
		if existingUser != nil && existingUser.ID != id {
			return dto.UserResponse{}, constants.ErrEmailTaken
		}

	}
	if user.Password != "" {
		hashedPassword, err := auth.HashPassword(user.Password)
		if err != nil {
			return dto.UserResponse{}, err
		}
		user.Password = hashedPassword
	}

	err = s.Repo.Update(user)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Map user to response DTO
	userResponse := mappers.MapUserToResponse(user)
	return userResponse, nil
}

// Delete User
func (s *UserService) DeleteUser(id uint) error {
	_, err := s.Repo.GetByID(id, []string{"id"})
	if err != nil {
		return constants.ErrUserNotFound
	}

	return s.Repo.Delete(id)
}
