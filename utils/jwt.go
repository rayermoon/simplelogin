package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key to sign JWT tokens
var secretKey = []byte("secret")

// Claims struct to store custom claims in the JWT token
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT token with the provided user ID
func GenerateJWT(userID string) (string, error) {
	// Create a new set of claims
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create and sign the JWT token with the claims and secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyJWT verifies the JWT token and returns the claims if valid
func VerifyJWT(tokenString string) (*Claims, error) {
	// Parse the token with the secret key
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
