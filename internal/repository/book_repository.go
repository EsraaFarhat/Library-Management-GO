package repository

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

// Create Book
func (r *BookRepository) Create(book *models.Book) (*models.Book, error) {
	err := r.DB.Create(book).Error
	return book, err
}

// Get Book by ID
func (r *BookRepository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.DB.First(&book, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constants.ErrBookNotFound
	}
	return &book, err
}

// Get All Books
func (r *BookRepository) GetAll(page, limit int) ([]models.Book, int64, error) {
	var books []models.Book
	var total int64

	// Count total books
	r.DB.Model(&models.Book{}).Count(&total)

	// Pagination logic
	offset := (page - 1) * limit

	// Fetch books with pagination and sorting
	err := r.DB.Order("created_at DESC").Limit(limit).Offset(offset).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}

	return books, total, nil
}

// Get Book by ISBN
func (r *BookRepository) GetByISBN(isbn string) (*models.Book, error) {
	var book models.Book
	err := r.DB.Where("isbn = ?", isbn).First(&book).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constants.ErrBookNotFound
	}
	return &book, err
}

// Update Book
func (r *BookRepository) Update(book *models.Book) error {
	return r.DB.Save(book).Error
}

// Delete Book
func (r *BookRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Book{}, id).Error
}

func (r *BookRepository) DecreaseBookCopies(bookID uint) error {
    book, err := r.GetByID(bookID)
    if err != nil {
        return err
    }
    // if book.CopiesAvailable < copies {
    //     return errors.New("not enough copies available")
    // }
    book.CopiesAvailable -= 1
    return r.DB.Save(&book).Error
}

func (r *BookRepository) IncreaseBookCopies(bookID uint) error {
    book, err := r.GetByID(bookID)
    if err != nil {
        return err
    }
    book.CopiesAvailable += 1
    return r.DB.Save(&book).Error
}
