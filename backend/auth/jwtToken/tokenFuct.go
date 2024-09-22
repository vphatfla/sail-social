package jwtToken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Define a secret key for signing tokens
var secretKey = []byte("your-secret-key")

func GenerateJWTToken(userId int64, typeUser string) (string, error) {
	// claims
	claims := jwt.MapClaims{
		"userId":   userId,
		"typeUser": typeUser,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyJWTToken(tokenString string, userId string, typeUser string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	// Check if the token is valid
	if err != nil || !token.Valid {
		return false, err
	}

	// Extract and verify claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the user_id claim matches the provided userID
		if claims["user_id"] == userId && claims["typeUser"] == typeUser {
			return true, nil
		} else {
			return false, fmt.Errorf("user ID or/and type user does not match")
		}
	}

	return false, fmt.Errorf("invalid token claims")
}
