package repository

import (
	"errors"

	"library-management/internal/models"

	"gorm.io/gorm"
)

type BorrowRepository struct {
	DB *gorm.DB
}

func NewBorrowRepository(db *gorm.DB) *BorrowRepository {
	return &BorrowRepository{DB: db}
}

// Create a new borrow record
func (r *BorrowRepository) Create(borrow *models.Borrow) error {
	return r.DB.Create(borrow).Error
}

// Get All Borrows
func (r *BorrowRepository) GetAll(page, limit int) ([]models.Borrow, int64, error) {
	var borrows []models.Borrow
	var total int64

	// Count total borrows
	r.DB.Model(&models.Borrow{}).Count(&total)

	// Pagination logic
	offset := (page - 1) * limit

	// Fetch borrows with pagination and sorting
	err := r.DB.Order("created_at DESC").Limit(limit).Offset(offset).Preload("Book").
		Preload("User").Find(&borrows).Error
	if err != nil {
		return nil, 0, err
	}

	return borrows, total, nil
}

// Get a borrow record by UserID and BookID
func (r *BorrowRepository) GetBorrowRecord(userID, bookID uint) (*models.Borrow, error) {
	var borrow models.Borrow
	err := r.DB.Where("user_id = ? AND book_id = ?", userID, bookID).First(&borrow).Error
	if err != nil {
		return nil, errors.New("borrow record not found")
	}
	return &borrow, nil
}

// Delete a borrow record when a book is returned
func (r *BorrowRepository) Delete(borrow *models.Borrow) error {
	return r.DB.Delete(borrow).Error
}

// GetBorrowsByUserID retrieves borrow records for a specific user
func (r *BorrowRepository) GetBorrowsByUserID(userID uint, page, limit int) ([]models.Borrow, int64, error) {
	var borrows []models.Borrow
	var total int64

	// Count total borrows for the user
	if err := r.DB.Model(&models.Borrow{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Fetch borrows with pagination and preload book and user details
	offset := (page - 1) * limit
	err := r.DB.Where("user_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Preload("Book").
		// Preload("User").
		Find(&borrows).Error
	if err != nil {
		return nil, 0, err
	}

	return borrows, total, nil
}
