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
)

// User Errors
var (
	ErrInvalidUserID   = errors.New("invalid user id")
	ErrUserNotFound    = errors.New("user not found")
	ErrEmailTaken      = errors.New("email is already registered")
	ErrInvalidUserRole = errors.New("invalid user role")
)

// Book Errors
var (
	ErrInvalidBookID       = errors.New("invalid book id")
	ErrBookNotFound        = errors.New("book not found")
	ErrBookAlreadyBorrowed = errors.New("book is already borrowed")
	ErrBookLimitExceeded   = errors.New("borrow limit exceeded")
	ErrNoCopiesAvailable   = errors.New("no copies available")
	ErrISBNExists          = errors.New("isbn is already registered")
)

// Validation Errors
var (
	ErrInvalidInput = errors.New("invalid input data")
	ErrDeleteFailed = errors.New("failed to delete")
)

// Server Errors
var (
	ErrInternalServer = errors.New("internal server error")
)
