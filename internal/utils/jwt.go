package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Custom claims structure
type Claims struct {
	UserID uint   `json:"userID"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Generate JWT Token
func GenerateToken(userID uint, role string) (string, error) {
	jwtSecret := []byte(os.Getenv("SECRET_KEY")) // Fetch dynamically

	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
