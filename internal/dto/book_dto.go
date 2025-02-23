package dto

import "time"

// BookCreateRequest represents the input for book creation.
type BookCreateRequest struct {
	Title           string    `json:"title" validate:"required"`
	Author          string    `json:"author" validate:"required"`
	ISBN            string    `json:"isbn" validate:"required"`
	CopiesAvailable int       `json:"copies_available" validate:"required,min=1"`
	PublishedAt     time.Time `json:"published_at" validate:"required"`
}

// BookUpdateRequest represents the input for book update.
type BookUpdateRequest struct {
	Title           *string    `json:"title,omitempty"`
	Author          *string    `json:"author,omitempty"`
	ISBN            *string    `json:"isbn,omitempty"`
	CopiesAvailable *int       `json:"copies_available,omitempty" validate:"omitempty,min=1"`
	PublishedAt     *time.Time `json:"published_at,omitempty"`
}

// BookResponse represents the output for book-related endpoints.
type BookResponse struct {
	ID              uint      `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	ISBN            string    `json:"isbn"`
	CopiesAvailable int       `json:"copies_available"`
	PublishedAt     time.Time `json:"published_at"`
}
