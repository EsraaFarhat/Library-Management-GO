package services

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/models"
	"library-management/internal/repository"
)

type BorrowServiceInterface interface {
	BorrowBook(req dto.BorrowCreateRequest, userIDUint uint) error
	ReturnBook(req dto.ReturnRequest, userIDUint uint) error
	GetBorrowRecords(page, limit int) ([]dto.BorrowResponse, int64, error)
	GetUserBorrows(userID uint, page, limit int) ([]dto.BorrowResponse, int64, error)
}

type BorrowService struct {
	BorrowRepo repository.BorrowRepositoryInterface
	BookRepo   repository.BookRepositoryInterface
	UserRepo   repository.UserRepositoryInterface
}

func NewBorrowService(borrowRepo repository.BorrowRepositoryInterface, bookRepo repository.BookRepositoryInterface, userRepo repository.UserRepositoryInterface) BorrowServiceInterface {
	return &BorrowService{
		BorrowRepo: borrowRepo,
		BookRepo:   bookRepo,
		UserRepo:   userRepo,
	}
}

// BorrowBook handles borrowing a book
func (s *BorrowService) BorrowBook(req dto.BorrowCreateRequest, userIDUint uint) error {
	// Check if the book exists
	book, err := s.BookRepo.GetByID(req.BookID, nil)
	if err != nil {
		return constants.ErrBookNotFound
	}

	// Ensure the book has available copies
	if book.CopiesAvailable <= 0 {
		return constants.ErrBookNotAvailable
	}

	// Check if the user has already borrowed this book
	// existingBorrow, _ := s.BorrowRepo.GetBorrowRecord(req.UserID, req.BookID)
	// if existingBorrow != nil {
	// 	return nil, errors.New("user has already borrowed this book")
	// }

	// Create the borrow record
	borrow := &models.Borrow{
		UserID:  userIDUint,
		BookID:  req.BookID,
		DueDate: req.DueDate,
	}

	// Start transaction
	tx, borrowRepo := s.BorrowRepo.BeginTransaction()

	if err := s.BorrowRepo.Create(borrow); err != nil {
		borrowRepo.RollbackTransaction(tx)
		return err
	}

	if err := s.BookRepo.DecreaseBookCopies(req.BookID); err != nil {
		borrowRepo.RollbackTransaction(tx)
		return err
	}

	borrowRepo.CommitTransaction(tx)
	return nil
}

// ReturnBook handles returning a borrowed book
func (s *BorrowService) ReturnBook(req dto.ReturnRequest, userIDUint uint) error {

	// Check if borrow record exists
	borrow, err := s.BorrowRepo.GetBorrowRecord(userIDUint, req.BookID)
	if err != nil {
		return constants.ErrBorrowNotFound
	}

	// Start transaction
	tx, borrowRepo := s.BorrowRepo.BeginTransaction()

	// Delete the borrow record
	if err := s.BorrowRepo.Delete(borrow); err != nil {
		borrowRepo.RollbackTransaction(tx)
		return err
	}

	// Increase book copies only if it was borrowed
	if err := s.BookRepo.IncreaseBookCopies(req.BookID); err != nil {
		borrowRepo.RollbackTransaction(tx)
		return err
	}

	borrowRepo.CommitTransaction(tx)
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
