package repository

import (
	"errors"
	"library-management/internal/constants"
	"library-management/internal/models"

	"gorm.io/gorm"
)

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
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constants.ErrUserNotFound
	}
	return &user, err
}

// Get User by Email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
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
