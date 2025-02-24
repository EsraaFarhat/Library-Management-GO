package constants

import "errors"

// Authentication Errors
var (
	ErrInvalidCredentials    = errors.New("invalid email or password")
	ErrUnauthorized          = errors.New("unauthorized access")
	ErrForbidden             = errors.New("forbidden: insufficient permissions")
	ErrMissingAuthHeader     = errors.New("authorization header missing")
	ErrInvalidTokenFormat    = errors.New("invalid token format")
	ErrInvalidOrExpiredToken = errors.New("invalid or expired token")
	ErrInvalidSigningMethod  = errors.New("unexpected signing method")
)

// User Errors
var (
	ErrInvalidUserID = errors.New("invalid user id")
	ErrUserNotFound  = errors.New("user not found")
	ErrEmailTaken    = errors.New("email is already registered")
)

// Book Errors
var (
	ErrInvalidBookID = errors.New("invalid book id")
	ErrBookNotFound  = errors.New("book not found")
	ErrISBNExists    = errors.New("isbn is already registered")
)

// Borrow Errors
var (
	ErrBorrowNotFound   = errors.New("borrow not found")
	ErrBookNotAvailable = errors.New("book is not available for borrowing")
)

// Validation Errors
var (
	ErrInvalidInput = errors.New("invalid input data")
)

// Server Errors
var (
	ErrInternalServer = errors.New("internal server error")
)
