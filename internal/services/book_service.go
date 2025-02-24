package services

import (
	"library-management/internal/constants"
	"library-management/internal/dto"
	"library-management/internal/repository"
	"library-management/internal/utils/mappers"
)

type BookServiceInterface interface {
	CreateBook(req dto.BookCreateRequest) (dto.BookResponse, error)
	GetBook(id uint, fields []string) (dto.BookResponse, error)
	GetAllBooks(page, limit int, fields []string) ([]dto.BookResponse, int64, error)
	UpdateBook(id uint, req dto.BookUpdateRequest) (dto.BookResponse, error)
	DeleteBook(id uint) error
}

type BookService struct {
	Repo repository.BookRepositoryInterface
}

func NewBookService(repo repository.BookRepositoryInterface) BookServiceInterface {
	return &BookService{Repo: repo}
}

// Create Book
func (s *BookService) CreateBook(req dto.BookCreateRequest) (dto.BookResponse, error) {
	book := mappers.MapCreateRequestToBook(req)

	// Check if the ISBN already exists
	existingBook, _ := s.Repo.GetByISBN(book.ISBN)
	if existingBook != nil {
		return dto.BookResponse{}, constants.ErrISBNExists
	}

	// Save book in the database
	book, err := s.Repo.Create(book)
	if err != nil {
		return dto.BookResponse{}, err
	}
	// Map book to response DTO
	bookResponse := mappers.MapBookToResponse(book)
	return bookResponse, nil
}

// Get Book by ID
func (s *BookService) GetBook(id uint, fields []string) (dto.BookResponse, error) {
	book, err := s.Repo.GetByID(id, fields)
	if err != nil {
		return dto.BookResponse{}, err
	}

	// Map book to response DTO
	bookResponse := mappers.MapBookToResponse(book)
	return bookResponse, nil
}

// Get All Books
func (s *BookService) GetAllBooks(page, limit int, fields []string) ([]dto.BookResponse, int64, error) {
	// Fetch books from the repository
	books, total, err := s.Repo.GetAll(page, limit, fields)
	if err != nil {
		return nil, 0, err
	}

	// Map each models.Book to dto.BookResponse
	bookResponses := make([]dto.BookResponse, len(books))
	for i, book := range books {
		bookResponses[i] = mappers.MapBookToResponse(&book)
	}

	return bookResponses, total, nil
}

// Update Book
func (s *BookService) UpdateBook(id uint, req dto.BookUpdateRequest) (dto.BookResponse, error) {
	book, err := s.Repo.GetByID(id, []string{})
	if err != nil {
		return dto.BookResponse{}, constants.ErrBookNotFound
	}

	// Update book fields
	mappers.UpdateBookFromDTO(book, req)

	if req.ISBN != nil {
		// Check if the isbn is already in use by another book
		existingBook, _ := s.Repo.GetByISBN(book.ISBN)
		if existingBook != nil && existingBook.ID != id {
			return dto.BookResponse{}, constants.ErrISBNExists
		}

	}

	err = s.Repo.Update(book)
	if err != nil {
		return dto.BookResponse{}, err
	}

	// Map book to response DTO
	bookResponse := mappers.MapBookToResponse(book)
	return bookResponse, nil
}

// Delete Book
func (s *BookService) DeleteBook(id uint) error {
	_, err := s.Repo.GetByID(id, []string{})
	if err != nil {
		return constants.ErrBookNotFound
	}

	return s.Repo.Delete(id)
}
