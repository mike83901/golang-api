package utils

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("90505e38-0616-434f-9d8d-5ee99acd42a8") // Replace with your JWT secret key

// Claims defines the structure of the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// ExtractUsernameFromToken extracts the username from a JWT token string
func ExtractUsernameFromToken(tokenString string) (string, error) {
	// Check if the token string has a "Bearer " prefix and remove it
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	// Check if the token is valid and type-cast claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", errors.New("invalid token")
}
