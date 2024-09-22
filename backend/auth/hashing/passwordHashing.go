package hashing

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwordStr string) ([]byte, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(passwordStr), bcrypt.DefaultCost)
	return b, err
}
