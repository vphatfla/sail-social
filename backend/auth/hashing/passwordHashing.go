package hashing

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwordStr string) ([]byte, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(passwordStr), bcrypt.DefaultCost)
	return b, err
}

func VerifyPassword(hashedPassword []byte, password string) error {
	bPW := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashedPassword, bPW)
	return err
}
