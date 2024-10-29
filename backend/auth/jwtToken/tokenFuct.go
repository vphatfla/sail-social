package jwtToken

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Define a secret key for signing tokens
var secretKey = []byte("your-secret-key")

func GenerateJWTToken(userId int, typeUser string, is_onboarded bool) (string, error) {
	// claims
	claims := jwt.MapClaims{
		"userId":      strconv.Itoa(int(userId)), // convert to string
		"typeUser":    typeUser,
		"isOnboarded": is_onboarded,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyJWTToken(tokenString string) (int, string, error) {
	var id int
	var typeUser string

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	// Check if the token is valid
	if err != nil {
		return id, typeUser, err
	}

	if !token.Valid {
		return id, typeUser, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimUserId := claims["userId"]
		if intValue, ok := claimUserId.(string); ok {
			id, err = strconv.Atoi(intValue)
			if err != nil {
				return id, typeUser, fmt.Errorf("invalid token userId")
			}
		} else {
			return id, typeUser, fmt.Errorf("invalid token userId")
		}

		claimUserType := claims["typeUser"]
		if strValue, ok := claimUserType.(string); ok {
			typeUser = strValue
		} else {
			return id, typeUser, fmt.Errorf("invalid token typeUser")
		}

	}

	return id, typeUser, nil
}
