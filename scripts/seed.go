package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Book struct {
	ID              uint      `gorm:"primaryKey"`
	Title           string    `gorm:"not null"`
	Author          string    `gorm:"not null"`
	ISBN            string    `gorm:"unique;not null"`
	CopiesAvailable int       `gorm:"not null"`
	PublishedAt     time.Time `gorm:"not null"`
}

// HashPassword hashes a given password
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func main() {
	// Load .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }
	// Load environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Database connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Auto-migrate tables
	err = db.AutoMigrate(&User{}, &Book{})
	if err != nil {
		log.Fatalf("❌ Failed to auto-migrate: %v", err)
	}

	seedUsers(db)
	seedBooks(db)
}

// Seed Users
func seedUsers(db *gorm.DB) {
	var count int64
	db.Model(&User{}).Count(&count)

	if count > 0 {
		fmt.Println("⚠️ Users table already has data, skipping seeding.")
		return
	}

	hashedPassword, err := HashPassword("Aa12345@")
	if err != nil {
		log.Fatalf("❌ Error hashing password: %v", err)
	}

	users := []User{
		{Name: "Admin User", Email: "admin@example.com", Password: hashedPassword, Role: "admin"},
		{Name: "Regular User", Email: "user@example.com", Password: hashedPassword, Role: "user"},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatalf("❌ Failed to seed users: %v", err)
	}
	fmt.Println("✅ Seeded users successfully!")
}

// Seed Books
func seedBooks(db *gorm.DB) {
	var count int64
	db.Model(&Book{}).Count(&count)

	if count > 0 {
		fmt.Println("⚠️ Books table already has data, skipping seeding.")
		return
	}

	books := []Book{
		{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", ISBN: "978-0134190440", CopiesAvailable: 5, PublishedAt: time.Now()},
		{Title: "Clean Code", Author: "Robert C. Martin", ISBN: "978-0132350884", CopiesAvailable: 3, PublishedAt: time.Now()},
	}

	if err := db.Create(&books).Error; err != nil {
		log.Fatalf("❌ Failed to seed books: %v", err)
	}
	fmt.Println("✅ Seeded books successfully!")
}
