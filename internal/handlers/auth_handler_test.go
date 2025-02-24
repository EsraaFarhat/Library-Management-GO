package handlers_test

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/handlers"
	"library-management/internal/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.AuthServiceInterface)
	handler := handlers.NewAuthHandler(mockService) // reqBody := `{"email": "test@example.com", "password": "Aa12345@"}`

	reqBody := `{"name": "test", "email": "test@example.com", "password": "Aa12345@"}`
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Define expected user response
	expectedUser := dto.UserResponse{
		ID:    1,
		Name:  "test user",
		Email: "test@example.com",
		Role:  "member",
	}

	// Mock the service call
	mockService.On("Register", mock.Anything).Return("mockToken", expectedUser, nil)
	handler.Register(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestRegisterHandler_EmailTaken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.AuthServiceInterface)
	handler := handlers.NewAuthHandler(mockService)
	reqBody := `{"name": "test", "email": "test@example.com", "password": "Aa12345@"}`
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	mockService.On("Register", mock.Anything).Return("", dto.UserResponse{}, constants.ErrEmailTaken)
	handler.Register(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	mockService.AssertExpectations(t)
}

func TestLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.AuthServiceInterface)
	handler := handlers.NewAuthHandler(mockService)
	reqBody := `{"email": "test@example.com", "password": "Aa12345@"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Define expected user response
	expectedUser := dto.UserResponse{
		ID:    1,
		Name:  "test user",
		Email: "test@example.com",
		Role:  "member",
	}

	mockService.On("Login", mock.Anything).Return("mockToken", expectedUser, nil)
	handler.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.AuthServiceInterface)
	handler := handlers.NewAuthHandler(mockService)
	reqBody := `{"email": "test@example.com", "password": "Bb12789@"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	mockService.On("Login", mock.Anything).Return("", dto.UserResponse{}, constants.ErrInvalidCredentials)
	handler.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockService.AssertExpectations(t)
}
