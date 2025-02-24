package handlers

import (
	"library-management/internal/dto"
	"library-management/internal/services"
	"library-management/internal/utils/error_handlers"
	"library-management/internal/utils/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service services.AuthServiceInterface
}

func NewAuthHandler(service services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{Service: service}
}

// Register a new user
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.UserRegisterRequest
	if err := handlers.BindAndValidate(c, &req); err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, err)
		return
	}
	token, user, err := h.Service.Register(req)
	if err != nil {
		error_handlers.HandleAuthError(c, err)
		return
	}
	handlers.RespondWithSuccess(c, http.StatusCreated, map[string]interface{}{"token": token, "user": user})
}

// Login User
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest

	if err := handlers.BindAndValidate(c, &req); err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	token, user, err := h.Service.Login(req)
	if err != nil {
		error_handlers.HandleAuthError(c, err)
		return
	}

	handlers.RespondWithSuccess(c, http.StatusOK, map[string]interface{}{"token": token, "user": user})
}
