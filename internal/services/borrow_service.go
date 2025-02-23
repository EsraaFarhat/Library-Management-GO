package services

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/models"
	"library-management/internal/repository"
	"library-management/internal/utils"
	"log"
	"github.com/go-playground/validator/v10"
	"errors"
)

var borrowValidator = validator.New()

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
	// Validate struct
	if err := borrowValidator.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedError := utils.FormatValidationErrors(validationErrors, &req)
			return nil, formattedError
		}
		return nil, err
	}

	// Check if the user exists
	_, err := s.UserRepo.GetByID(req.UserID)
	if err != nil {
		return nil, constants.ErrUserNotFound
	}

	// Check if the book exists
	book, err := s.BookRepo.GetByID(req.BookID)
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

	if err := s.BorrowRepo.CreateBorrowRecord(borrow); err != nil {
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
	// Validate struct
	if err := borrowValidator.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedError := utils.FormatValidationErrors(validationErrors, &req)
			return formattedError
		}
		return err
	}
	log.Println(req)

	// Check if borrow record exists
	borrow, err := s.BorrowRepo.GetBorrowRecord(req.UserID, req.BookID)
	if err != nil {
		log.Println(err)
		return constants.ErrBorrowNotFound
	}

	// Start transaction
	tx := s.BorrowRepo.DB.Begin()

	// Delete the borrow record
	if err := s.BorrowRepo.DeleteBorrowRecord(borrow); err != nil {
		log.Println(err)
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
func (s *BorrowService) GetBorrowRecords(page, limit int) ([]models.Borrow, int64, error) {
	return s.BorrowRepo.GetAll(page, limit)
}

// GetUserBorrows retrieves borrow records for a specific user
func (s *BorrowService) GetUserBorrows(userID uint) ([]models.Borrow, error) {
	return s.BorrowRepo.GetBorrowsByUserID(userID)
}