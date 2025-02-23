package services

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/models"
	"library-management/internal/repository"
)

type BorrowService struct {
	BorrowRepo *repository.BorrowRepository
	BookRepo   *repository.BookRepository
	UserRepo   *repository.UserRepository
}

func NewBorrowService(borrowRepo *repository.BorrowRepository, bookRepo *repository.BookRepository, userRepo *repository.UserRepository) *BorrowService {
	return &BorrowService{
		BorrowRepo: borrowRepo,
		BookRepo:   bookRepo,
		UserRepo:   userRepo,
	}
}

// BorrowBook handles borrowing a book
func (s *BorrowService) BorrowBook(req dto.BorrowCreateRequest) (*models.Borrow, error) {

	// Check if the user exists
	_, err := s.UserRepo.GetByID(req.UserID, []string{"id"})
	if err != nil {
		return nil, constants.ErrUserNotFound
	}

	// Check if the book exists
	book, err := s.BookRepo.GetByID(req.BookID, nil)
	if err != nil {
		return nil, constants.ErrBookNotFound
	}

	// Ensure the book has available copies
	if book.CopiesAvailable <= 0 {
		return nil, errors.New("book is not available for borrowing")
	}

	// Check if the user has already borrowed this book
	// existingBorrow, _ := s.BorrowRepo.GetBorrowRecord(req.UserID, req.BookID)
	// if existingBorrow != nil {
	// 	return nil, errors.New("user has already borrowed this book")
	// }

	// Create the borrow record
	borrow := &models.Borrow{
		UserID:  req.UserID,
		BookID:  req.BookID,
		DueDate: req.DueDate,
	}

	// Start transaction
	tx := s.BorrowRepo.DB.Begin()

	if err := s.BorrowRepo.Create(borrow); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := s.BookRepo.DecreaseBookCopies(req.BookID); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return borrow, nil
}

// ReturnBook handles returning a borrowed book
func (s *BorrowService) ReturnBook(req dto.ReturnRequest) error {

	// Check if borrow record exists
	borrow, err := s.BorrowRepo.GetBorrowRecord(req.UserID, req.BookID)
	if err != nil {
		return constants.ErrBorrowNotFound
	}

	// Start transaction
	tx := s.BorrowRepo.DB.Begin()

	// Delete the borrow record
	if err := s.BorrowRepo.Delete(borrow); err != nil {
		tx.Rollback()
		return err
	}

	// Increase book copies only if it was borrowed
	if err := s.BookRepo.IncreaseBookCopies(req.BookID); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// GetBorrowRecords retrieves all borrow records with pagination
func (s *BorrowService) GetBorrowRecords(page, limit int) ([]dto.BorrowResponse, int64, error) {
	// Fetch borrows from the repository
	borrows, total, err := s.BorrowRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	// Map each models.Borrow to dto.BorrowResponse
	borrowResponses := make([]dto.BorrowResponse, len(borrows))
	for i, borrow := range borrows {
		borrowResponses[i] = dto.BorrowResponse{
			ID:      borrow.ID,
			DueDate: borrow.DueDate,
			User: &dto.UserResponse{
				ID:    borrow.User.ID,
				Name:  borrow.User.Name,
				Email: borrow.User.Email,
				Role:  borrow.User.Role,
			},
			Book: &dto.BookResponse{
				ID:              borrow.Book.ID,
				Title:           borrow.Book.Title,
				Author:          borrow.Book.Author,
				ISBN:            borrow.Book.ISBN,
				CopiesAvailable: borrow.Book.CopiesAvailable,
				PublishedAt:     borrow.Book.PublishedAt,
			},
		}
	}

	return borrowResponses, total, nil
}

// GetUserBorrows retrieves borrow records for a specific user
func (s *BorrowService) GetUserBorrows(userID uint, page, limit int) ([]dto.BorrowResponse, int64, error) {
	// Fetch borrows from the repository
	borrows, total, err := s.BorrowRepo.GetBorrowsByUserID(userID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	// Map each models.Borrow to dto.BorrowResponse
	borrowResponses := make([]dto.BorrowResponse, len(borrows))
	for i, borrow := range borrows {
		borrowResponses[i] = dto.BorrowResponse{
			ID:      borrow.ID,
			DueDate: borrow.DueDate,
			Book: &dto.BookResponse{
				ID:              borrow.Book.ID,
				Title:           borrow.Book.Title,
				Author:          borrow.Book.Author,
				ISBN:            borrow.Book.ISBN,
				CopiesAvailable: borrow.Book.CopiesAvailable,
				PublishedAt:     borrow.Book.PublishedAt,
			},
		}
	}

	return borrowResponses, total, nil
}
