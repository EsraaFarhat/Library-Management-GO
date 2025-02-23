package mappers

import (
	"library-management/internal/dto"
	"library-management/internal/models"
)

// MapBookToResponse maps a models.Book to a BookResponse.
func MapBookToResponse(book *models.Book) dto.BookResponse {
	return dto.BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		Author:          book.Author,
		ISBN:            book.ISBN,
		CopiesAvailable: book.CopiesAvailable,
		PublishedAt:     book.PublishedAt,
	}
}
