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

type BorrowHandler struct {
	Service services.BorrowServiceInterface
}

func NewBorrowHandler(service services.BorrowServiceInterface) *BorrowHandler {
	return &BorrowHandler{Service: service}
}

// BorrowBook handles borrowing a book
func (h *BorrowHandler) BorrowBook(c *gin.Context) {
	var req dto.BorrowCreateRequest
	if err := handlers.BindAndValidate(c, &req); err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	borrow, err := h.Service.BorrowBook(req)
	if err != nil {
		error_handlers.HandleBorrowError(c, err)
		return
	}
	handlers.RespondWithSuccess(c, http.StatusCreated, borrow)
}

// ReturnBook handles returning a borrowed book
func (h *BorrowHandler) ReturnBook(c *gin.Context) {
	var req dto.ReturnRequest
	if err := handlers.BindAndValidate(c, &req); err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	err := h.Service.ReturnBook(req)
	if err != nil {
		error_handlers.HandleBorrowError(c, err)
		return
	}
	handlers.RespondWithSuccess(c, http.StatusOK, "Book returned successfully")
}

// GetBorrowRecords retrieves all borrow records with pagination
func (h *BorrowHandler) GetBorrowRecords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	records, total, err := h.Service.GetBorrowRecords(page, limit)
	if err != nil {
		error_handlers.HandleBorrowError(c, err)
		return
	}

	response := map[string]interface{}{
		"rows":  records,
		"total": total,
		"page":  page,
		"limit": limit,
	}
	handlers.RespondWithSuccess(c, http.StatusOK, response)
}

// GetUserBorrows retrieves borrow records for a specific user
func (h *BorrowHandler) GetUserBorrows(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidUserID)
		return
	}

	// Parse pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	borrows, total, err := h.Service.GetUserBorrows(uint(userID), page, limit)
	if err != nil {
		error_handlers.HandleBorrowError(c, err)
		return
	}

	// Respond with pagination metadata
	response := map[string]interface{}{
		"rows":  borrows,
		"total": total,
		"page":  page,
		"limit": limit,
	}
	handlers.RespondWithSuccess(c, http.StatusOK, response)
}
