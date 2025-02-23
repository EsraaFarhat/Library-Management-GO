package services

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/models"
	"library-management/internal/repository"
	"library-management/internal/utils"

	"github.com/go-playground/validator/v10"
)

var bookValidator = validator.New()

type BookService struct {
	Repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{Repo: repo}
}

// Create Book
func (s *BookService) CreateBook(req dto.BookCreateRequest) (*models.Book, error) {
	book := utils.MapCreateRequestToBook(req)

	// Validate struct
	if err := bookValidator.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return nil, utils.FormatValidationErrors(validationErrors)
		}
		return nil, err
	}

	// Check if the ISBN already exists
	existingBook, _ := s.Repo.GetByISBN(book.ISBN)
	if existingBook != nil {
		return nil, constants.ErrISBNExists
	}

	// Save book in the database
	book, err := s.Repo.Create(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

// Get Book by ID
func (s *BookService) GetBook(id uint) (*models.Book, error) {
	return s.Repo.GetByID(id)
}

// Get All Books
func (s *BookService) GetAllBooks(page, limit int) ([]models.Book, int64, error) {
	return s.Repo.GetAll(page, limit)
}

// Update Book
func (s *BookService) UpdateBook(id uint, req dto.BookUpdateRequest) (*models.Book, error) {
	book, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, constants.ErrBookNotFound
	}

	// Update book fields
	utils.UpdateBookFromDTO(book, req)

	// Validate struct
	if err := bookValidator.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return nil, utils.FormatValidationErrors(validationErrors)
		}
		return nil, err
	}

	if req.ISBN != nil {
		// Check if the isbn is already in use by another book
		existingBook, _ := s.Repo.GetByISBN(book.ISBN)
		if existingBook != nil && existingBook.ID != id {
			return nil, constants.ErrISBNExists
		}

	}

	err = s.Repo.Update(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// Delete Book
func (s *BookService) DeleteBook(id uint) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return constants.ErrBookNotFound
	}

	return s.Repo.Delete(id)
}
