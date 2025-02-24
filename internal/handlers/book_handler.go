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

type BookHandler struct {
	Service services.BookServiceInterface
}

func NewBookHandler(service services.BookServiceInterface) *BookHandler {
	return &BookHandler{Service: service}
}

// Create a new book
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.BookCreateRequest
	if err := handlers.BindAndValidate(c, &req); err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	createdBook, err := h.Service.CreateBook(req)
	if err != nil {
		error_handlers.HandleBookError(c, err)
		return
	}
	handlers.RespondWithSuccess(c, http.StatusCreated, createdBook)
}

// Get all books
func (h *BookHandler) GetAllBooks(c *gin.Context) {
	// Default values
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Fetch paginated books
	books, total, err := h.Service.GetAllBooks(page, limit, nil)
	if err != nil {
		error_handlers.HandleBookError(c, err)
		return
	}

	// Respond with pagination metadata
	response := map[string]interface{}{
		"rows":  books,
		"total": total,
		"page":  page,
		"limit": limit,
	}

	handlers.RespondWithSuccess(c, http.StatusOK, response)
}

// Get Book by ID
func (h *BookHandler) GetBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidBookID)
		return
	}

	book, err := h.Service.GetBook(uint(id), []string{})
	if err != nil {
		handlers.RespondWithError(c, http.StatusNotFound, constants.ErrBookNotFound)
		return
	}
	handlers.RespondWithSuccess(c, http.StatusOK, book)
}

// Update Book
func (h *BookHandler) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidBookID)
		return
	}

	var req dto.BookUpdateRequest
	if err := handlers.BindAndValidate(c, &req); err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	book, err := h.Service.UpdateBook(uint(id), req)
	if err != nil {
		error_handlers.HandleBookError(c, err)
		return
	}
	handlers.RespondWithSuccess(c, http.StatusOK, book)
}

// Delete Book
func (h *BookHandler) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handlers.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidBookID)
		return
	}

	err = h.Service.DeleteBook(uint(id))
	if err != nil {
		error_handlers.HandleBookError(c, err)
		return
	}
	handlers.RespondWithSuccess(c, http.StatusOK, map[string]interface{}{"id": id})
}
