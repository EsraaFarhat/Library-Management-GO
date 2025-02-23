package dto

import "time"

type BorrowCreateRequest struct {
	UserID  uint      `json:"user_id" validate:"required"`
	BookID  uint      `json:"book_id" validate:"required"`
	DueDate time.Time `json:"due_date" validate:"required"`
}

type ReturnRequest struct {
	UserID uint `json:"user_id" validate:"required"`
	BookID uint `json:"book_id" validate:"required"`
}
