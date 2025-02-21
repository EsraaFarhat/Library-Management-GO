package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
	Role     string `json:"role" gorm:"type:varchar(20);not null;default:'member'"`

	// A User can borrow many books
	Borrows []Borrow `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
