package services_test

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/mocks"
	"library-management/internal/models"
	"library-management/internal/services"
	"library-management/internal/utils/auth"
	"library-management/internal/utils/mappers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryInterface)
	authService := services.NewAuthService(mockRepo)

	req := dto.UserRegisterRequest{
		Name:     "John Doe",
		Email:    "johna@example.com",
		Password: "Aa12345@",
	}

	user := mappers.MapRegisterRequestToUser(req)
	user.Password, _ = auth.HashPassword(req.Password)
	mockRepo.On("GetByEmail", user.Email, mock.Anything).Return(nil, nil)
	mockRepo.On("Create", mock.Anything).Return(user, nil)

	token, userResponse, err := authService.Register(req)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, req.Email, userResponse.Email)
	mockRepo.AssertExpectations(t)
}

func TestRegister_EmailAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryInterface)
	authService := services.NewAuthService(mockRepo)

	req := dto.UserRegisterRequest{
		Email: "existing@example.com",
	}

	mockRepo.On("GetByEmail", req.Email, mock.Anything).Return(&models.User{}, nil)

	token, userResponse, err := authService.Register(req)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrEmailTaken, err)
	assert.Empty(t, token)
	assert.Empty(t, userResponse)
	mockRepo.AssertCalled(t, "GetByEmail", req.Email, mock.Anything)
	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryInterface)
	authService := services.NewAuthService(mockRepo)

	req := dto.UserLoginRequest{
		Email:    "johasn@example.com",
		Password: "Aa12345@",
	}

	hashedPassword, _ := auth.HashPassword(req.Password)
	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
	}

	mockRepo.On("GetByEmail", req.Email, mock.Anything).Return(&user, nil)

	token, userResponse, err := authService.Login(req)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, req.Email, userResponse.Email)
	mockRepo.AssertCalled(t, "GetByEmail", req.Email, mock.Anything)
	mockRepo.AssertExpectations(t)
}

func TestLogin_InvalidPassword(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryInterface)
	authService := services.NewAuthService(mockRepo)

	req := dto.UserLoginRequest{
		Email:    "john@example.com",
		Password: "Bb12789@",
	}

	hashedPassword, _ := auth.HashPassword("Aa12345@")
	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
	}

	mockRepo.On("GetByEmail", req.Email, mock.Anything).Return(&user, nil)

	token, userResponse, err := authService.Login(req)

	assert.Error(t, err)
	assert.Equal(t, constants.ErrInvalidCredentials, err)
	assert.Empty(t, token)
	assert.Empty(t, userResponse)
	mockRepo.AssertExpectations(t)
}
