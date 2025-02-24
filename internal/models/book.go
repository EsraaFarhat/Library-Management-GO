package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string    `json:"title" gorm:"type:varchar(200);not null"`
	Author          string    `json:"author" gorm:"type:varchar(100);not null"`
	ISBN            string    `json:"isbn" gorm:"type:varchar(20);not null"`
	CopiesAvailable int       `json:"copies_available" gorm:"not null;check:copies_available >= 0"`
	PublishedAt     time.Time `json:"published_at" gorm:"not null"`

	// A Book can be borrowed multiple times
	Borrows []Borrow `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;"`
}
