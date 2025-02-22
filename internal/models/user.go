package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null" validate:"required"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null" validate:"required,email"`
	Password string `json:"-" gorm:"type:varchar(255);not null" validate:"required"`
	Role     string `json:"role" gorm:"type:varchar(20);not null;default:'member'" validate:"required,oneof=admin member"`

	// A User can borrow many books
	Borrows []Borrow `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
