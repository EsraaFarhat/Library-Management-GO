package handlers

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/services"
	"library-management/internal/utils"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// Create a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidInput)
		return
	}

	createdUser, err := h.Service.CreateUser(req)

	if err != nil {
		// Handle known constant errors
		// Check if the error is a validation error
		var validationErr *utils.ValidationError
		if errors.As(err, &validationErr) {

			utils.RespondWithError(c, http.StatusBadRequest, validationErr) // Return validation errors
			return
		}

		if errors.Is(err, constants.ErrEmailTaken) {
			utils.RespondWithError(c, http.StatusConflict, constants.ErrEmailTaken)
			return
		}

		// Unexpected errors
		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, createdUser)
}

// Get User by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidUserID)
		return
	}

	user, err := h.Service.GetUser(uint(id))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, constants.ErrUserNotFound)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, user)
}

// Update User
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidUserID)
		return
	}

	var req dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidInput)
		return
	}

	user, err := h.Service.UpdateUser(uint(id), req)
	if err != nil {
		// Handle known constant errors
		// Check if the error is a validation error
		var validationErr *utils.ValidationError
		if errors.As(err, &validationErr) {

			utils.RespondWithError(c, http.StatusBadRequest, validationErr) // Return validation errors
			return
		}

		if errors.Is(err, constants.ErrUserNotFound) {
			utils.RespondWithError(c, http.StatusNotFound, constants.ErrUserNotFound)
			return
		}
		if errors.Is(err, constants.ErrInvalidInput) {
			utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}
		if errors.Is(err, constants.ErrEmailTaken) {
			utils.RespondWithError(c, http.StatusConflict, constants.ErrEmailTaken)
			return
		}

		// Unexpected errors
		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return

	}

	utils.RespondWithSuccess(c, http.StatusOK, user)
}

// Delete User
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidUserID)
		return
	}

	err = h.Service.DeleteUser(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrUserNotFound):
			utils.RespondWithError(c, http.StatusNotFound, constants.ErrUserNotFound)
		default:
			utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		}
		return

	}
	utils.RespondWithSuccess(c, http.StatusOK, map[string]interface{}{"id": id})
}
