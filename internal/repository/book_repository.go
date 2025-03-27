package repository

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/models"

	"gorm.io/gorm"
)

var defaultBookFields = []string{"id", "title", "author", "isbn", "copies_available", "published_at"}

type BookRepositoryInterface interface {
	Create(book *models.Book) (*models.Book, error)
	GetByID(id uint, fields []string) (*models.Book, error)
	GetAll(page, limit int, fields []string) ([]models.Book, int64, error)
	GetByISBN(isbn string) (*models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
	DecreaseBookCopies(bookID uint) error
	IncreaseBookCopies(bookID uint) error
}

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepositoryInterface {
	return &BookRepository{DB: db}
}

// Create Book
func (r *BookRepository) Create(book *models.Book) (*models.Book, error) {
	err := r.DB.Create(book).Error
	return book, err
}

// Get Book by ID
func (r *BookRepository) GetByID(id uint, fields []string) (*models.Book, error) {
	var book models.Book
	// Start with a base query
	query := r.DB.Model(&models.Book{})

	// Use default fields if no specific fields are provided
	if len(fields) == 0 {
		fields = defaultBookFields
	}
	// Select specific fields
	query = query.Select(fields)

	err := query.First(&book, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constants.ErrBookNotFound
	}
	return &book, err
}

// Get All Books
func (r *BookRepository) GetAll(page, limit int, fields []string) ([]models.Book, int64, error) {
	var books []models.Book
	var total int64

	// Start with a base query
	query := r.DB.Model(&models.Book{})

	// Use default fields if no specific fields are provided
	if len(fields) == 0 {
		fields = defaultBookFields
	}
	// Select specific fields
	query = query.Select(fields)

	// Count total books (without pagination)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination logic
	offset := (page - 1) * limit

	// Fetch books with pagination and sorting
	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&books).Error; err != nil {
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
	book, err := r.GetByID(bookID, nil)
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
	book, err := r.GetByID(bookID, nil)
	if err != nil {
		// do nothing (in case the book was deleted)
		return nil
	}
	book.CopiesAvailable += 1
	return r.DB.Save(&book).Error
}
