package auth

import (
	"library-management/internal/constants"
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

// ValidateToken verifies a JWT token and returns claims if valid
func ValidateToken(tokenString string) (*Claims, error) {
	jwtSecret := []byte(os.Getenv("SECRET_KEY")) // Fetch dynamically

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, constants.ErrInvalidSigningMethod
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, constants.ErrInvalidOrExpiredToken
}
