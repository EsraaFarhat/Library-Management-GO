package constants

import "errors"

// User Errors
var (
	ErrInvalidUserID   = errors.New("invalid user id")
	ErrUserNotFound    = errors.New("user not found")
	ErrEmailTaken      = errors.New("email is already registered")
	ErrInvalidUserRole = errors.New("invalid user role")
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
