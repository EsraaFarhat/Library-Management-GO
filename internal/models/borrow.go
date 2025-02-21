package models

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model
	UserID   uint      `json:"user_id" gorm:"not null"`
	BookID   uint      `json:"book_id" gorm:"not null"`
	DueDate  time.Time `json:"due_date" gorm:"not null"`
	Returned bool      `json:"returned" gorm:"not null;default:false"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Book Book `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;"`
}
