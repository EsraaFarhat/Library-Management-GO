package handlers

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/services"
	"library-management/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// Register a new user
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidInput)
		return
	}
	token, err := h.Service.Register(req)
	if err != nil {
		// Check if the error is a validation error
		var validationErr *utils.ValidationError
		if errors.As(err, &validationErr) {

			utils.RespondWithError(c, http.StatusBadRequest, validationErr) // Return validation errors
			return
		}

		// Handle known constant errors
		if errors.Is(err, constants.ErrEmailTaken) {
			utils.RespondWithError(c, http.StatusConflict, constants.ErrEmailTaken)
			return
		}

		// Unexpected errors
		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, map[string]interface{}{"token": token})
}

// Login User
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidInput)
		return
	}

	token, err := h.Service.Login(req)
	if err != nil {
		// Handle known constant errors
		// Check if the error is a validation error
		var validationErr *utils.ValidationError
		if errors.As(err, &validationErr) {

			utils.RespondWithError(c, http.StatusBadRequest, validationErr) // Return validation errors
			return
		}

		if errors.Is(err, constants.ErrInvalidCredentials) {
			utils.RespondWithError(c, http.StatusUnauthorized, constants.ErrInvalidCredentials)
			return
		}

		// Unexpected errors
		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, map[string]interface{}{"token": token})
}
