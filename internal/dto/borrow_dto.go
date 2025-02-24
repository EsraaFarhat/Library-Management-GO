package dto

import "time"

type BorrowCreateRequest struct {
	// UserID  uint      `json:"user_id" validate:"required"`
	BookID  uint      `json:"book_id" validate:"required"`
	DueDate time.Time `json:"due_date" validate:"required"`
}

type ReturnRequest struct {
	// UserID uint `json:"user_id" validate:"required"`
	BookID uint `json:"book_id" validate:"required"`
}

type BorrowResponse struct {
	ID      uint          `json:"id"`
	UserID  uint          `json:"user_id,omitempty"`
	BookID  uint          `json:"book_id,omitempty"`
	DueDate time.Time     `json:"due_date"`
	User    *UserResponse `json:"user,omitempty"` // Include user details
	Book    *BookResponse `json:"book,omitempty"` // Include book details
}
