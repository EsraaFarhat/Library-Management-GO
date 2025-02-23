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

type BorrowHandler struct {
	Service *services.BorrowService
}

func NewBorrowHandler(service *services.BorrowService) *BorrowHandler {
	return &BorrowHandler{Service: service}
}

// BorrowBook handles borrowing a book
func (h *BorrowHandler) BorrowBook(c *gin.Context) {
	var req dto.BorrowCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidInput)
		return
	}

	borrow, err := h.Service.BorrowBook(req)
	if err != nil {
		var validationErr *utils.ValidationError
		if errors.As(err, &validationErr) {
			utils.RespondWithError(c, http.StatusBadRequest, validationErr)
			return
		}
		if errors.Is(err, constants.ErrUserNotFound) {
			utils.RespondWithError(c, http.StatusBadRequest, constants.ErrUserNotFound)
			return
		}
		if errors.Is(err, constants.ErrBookNotFound) {
			utils.RespondWithError(c, http.StatusBadRequest, constants.ErrBookNotFound)
			return
		}

		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, borrow)
}

// ReturnBook handles returning a borrowed book
func (h *BorrowHandler) ReturnBook(c *gin.Context) {
	var req dto.ReturnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidInput)
		return
	}

	err := h.Service.ReturnBook(req)
	if err != nil {
		var validationErr *utils.ValidationError
		if errors.As(err, &validationErr) {
			utils.RespondWithError(c, http.StatusBadRequest, validationErr)
			return
		}
		if errors.Is(err, constants.ErrUserNotFound) {
			utils.RespondWithError(c, http.StatusBadRequest, constants.ErrUserNotFound)
			return
		}
		if errors.Is(err, constants.ErrBookNotFound) {
			utils.RespondWithError(c, http.StatusBadRequest, constants.ErrBookNotFound)
			return
		}
		if errors.Is(err, constants.ErrUserNotFound) {
			utils.RespondWithError(c, http.StatusBadRequest, constants.ErrUserNotFound)
			return
		}
		if errors.Is(err, constants.ErrBorrowNotFound) {
			utils.RespondWithError(c, http.StatusBadRequest, constants.ErrBorrowNotFound)
			return
		}

		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Book returned successfully")
}

// GetBorrowRecords retrieves all borrow records with pagination
func (h *BorrowHandler) GetBorrowRecords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	records, total, err := h.Service.GetBorrowRecords(page, limit)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return
	}

	response := map[string]interface{}{
		"rows":  records,
		"total": total,
		"page":  page,
		"limit": limit,
	}
	utils.RespondWithSuccess(c, http.StatusOK, response)
}

// GetUserBorrows retrieves borrow records for a specific user
func (h *BorrowHandler) GetUserBorrows(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidUserID)
		return
	}

	borrows, err := h.Service.GetUserBorrows(uint(userID))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, borrows)
}