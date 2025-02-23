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

type BookHandler struct {
	Service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

// Create a new book
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.BookCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidInput)
		return
	}

	createdBook, err := h.Service.CreateBook(req)
	if err != nil {
		var validationErr *utils.ValidationError
		if errors.As(err, &validationErr) {
			utils.RespondWithError(c, http.StatusBadRequest, validationErr)
			return
		}
		if errors.Is(err, constants.ErrISBNExists) {
			utils.RespondWithError(c, http.StatusConflict, constants.ErrISBNExists)
			return
		}

		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, createdBook)
}

// Get all books
func (h *BookHandler) GetAllBooks(c *gin.Context) {
	// Default values
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Fetch paginated books
	books, total, err := h.Service.GetAllBooks(page, limit)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return
	}

	// Respond with pagination metadata
	response := map[string]interface{}{
		"rows":  books,
		"total": total,
		"page":  page,
		"limit": limit,
	}

	utils.RespondWithSuccess(c, http.StatusOK, response)
}

// Get Book by ID
func (h *BookHandler) GetBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidBookID)
		return
	}

	book, err := h.Service.GetBook(uint(id))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, constants.ErrBookNotFound)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, book)
}

// Update Book
func (h *BookHandler) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidBookID)
		return
	}

	var req dto.BookUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidInput)
		return
	}

	book, err := h.Service.UpdateBook(uint(id), req)
	if err != nil {
		var validationErr *utils.ValidationError
		if errors.As(err, &validationErr) {
			utils.RespondWithError(c, http.StatusBadRequest, validationErr)
			return
		}
		if errors.Is(err, constants.ErrBookNotFound) {
			utils.RespondWithError(c, http.StatusNotFound, constants.ErrBookNotFound)
			return
		}
		if errors.Is(err, constants.ErrInvalidInput) {
			utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidInput)
			return
		}

		if errors.Is(err, constants.ErrISBNExists) {
			utils.RespondWithError(c, http.StatusConflict, constants.ErrISBNExists)
			return
		}

		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, book)
}

// Delete Book
func (h *BookHandler) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, constants.ErrInvalidBookID)
		return
	}

	err = h.Service.DeleteBook(uint(id))
	if err != nil {
		if errors.Is(err, constants.ErrBookNotFound) {
			utils.RespondWithError(c, http.StatusNotFound, constants.ErrBookNotFound)
		} else {
			utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrInternalServer)
		}
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, map[string]interface{}{"id": id})
}
