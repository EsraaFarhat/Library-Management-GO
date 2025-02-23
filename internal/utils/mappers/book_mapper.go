package mappers

import (
	"library-management/internal/dto"
	"library-management/internal/models"
)

// Convert BookCreateRequest to Book model
func MapCreateRequestToBook(req dto.BookCreateRequest) *models.Book {
	return &models.Book{
		Title:           req.Title,
		Author:          req.Author,
		ISBN:            req.ISBN,
		CopiesAvailable: req.CopiesAvailable,
		PublishedAt:     req.PublishedAt,
	}
}

// Update Book model from DTO
func UpdateBookFromDTO(book *models.Book, req dto.BookUpdateRequest) {
	if req.Title != nil {
		book.Title = *req.Title
	}
	if req.Author != nil {
		book.Author = *req.Author
	}
	if req.ISBN != nil {
		book.ISBN = *req.ISBN
	}
	if req.CopiesAvailable != nil {
		book.CopiesAvailable = *req.CopiesAvailable
	}
	if req.PublishedAt != nil {
		book.PublishedAt = *req.PublishedAt
	}
}
