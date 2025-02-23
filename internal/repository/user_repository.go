package repository

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/models"

	"gorm.io/gorm"
)

var defaultUserFields = []string{"id", "name", "email", "role", "created_at"}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create User
func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	err := r.DB.Create(user).Error
	return user, err
}

// Get User by ID
func (r *UserRepository) GetByID(id uint, fields []string) (*models.User, error) {
	var user models.User
	// Start with a base query
	query := r.DB.Model(&models.User{})

	// Use default fields if no specific fields are provided
	if len(fields) == 0 {
		fields = defaultUserFields
	}
	// Select specific fields
	query = query.Select(fields)

	err := query.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constants.ErrUserNotFound
	}
	return &user, err
}

// Get All Users
func (r *UserRepository) GetAll(page, limit int, fields []string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Start with a base query
	query := r.DB.Model(&models.User{})

	// Use default fields if no specific fields are provided
	if len(fields) == 0 {
		fields = defaultUserFields
	}
	// Select specific fields
	query = query.Select(fields)

	// Count total users (without pagination)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination logic
	offset := (page - 1) * limit

	// Fetch users with pagination and sorting
	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Get User by Email
func (r *UserRepository) GetByEmail(email string, fields []string) (*models.User, error) {
	var user models.User
	// Start with a base query
	query := r.DB.Model(&models.User{})

	// Use default fields if no specific fields are provided
	if len(fields) == 0 {
		fields = defaultUserFields
	}
	// Select specific fields
	query = query.Select(fields)

	err := query.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constants.ErrUserNotFound
	}
	return &user, err
}

// Update User
func (r *UserRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

// Delete User
func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
