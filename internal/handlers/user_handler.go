package handlers

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/services"
	"library-management/internal/utils/error_handlers"
	"library-management/internal/utils/handlers"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service services.UserServiceInterface
}

func NewUserHandler(service services.UserServiceInterface) *UserHandler {
	return &UserHandler{Service: service}
}

// Create a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.UserCreateRequest
	if err := handlers.BindAndValidate(c, &req); err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	createdUser, err := h.Service.CreateUser(req)

	if err != nil {
		error_handlers.HandleUserError(c, err)
		return
	}
	handlers.RespondWithSuccess(c, http.StatusCreated, createdUser)
}

// Get all users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	// Default values
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Fetch paginated users
	users, total, err := h.Service.GetAllUsers(page, limit, nil)
	if err != nil {
		error_handlers.HandleUserError(c, err)
		return
	}

	// Respond with pagination metadata
	response := map[string]interface{}{
		"rows":  users,
		"total": total,
		"page":  page,
		"limit": limit,
	}

	handlers.RespondWithSuccess(c, http.StatusOK, response)
}

// Get User by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidUserID)
		return
	}

	user, err := h.Service.GetUser(uint(id), []string{})
	if err != nil {
		handlers.RespondWithError(c, http.StatusNotFound, constants.ErrUserNotFound)
		return
	}
	handlers.RespondWithSuccess(c, http.StatusOK, user)
}

// Update User
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidUserID)
		return
	}

	var req dto.UserUpdateRequest
	if err := handlers.BindAndValidate(c, &req); err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.Service.UpdateUser(uint(id), req)
	if err != nil {
		error_handlers.HandleUserError(c, err)
		return
	}

	handlers.RespondWithSuccess(c, http.StatusOK, user)
}

// Delete User
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidUserID)
		return
	}

	err = h.Service.DeleteUser(uint(id))
	if err != nil {
		error_handlers.HandleUserError(c, err)
		return
	}
	handlers.RespondWithSuccess(c, http.StatusOK, map[string]interface{}{"id": id})
}
