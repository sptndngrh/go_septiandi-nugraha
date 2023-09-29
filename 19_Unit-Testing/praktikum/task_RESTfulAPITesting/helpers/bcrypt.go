package helpers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

// CheckPasswordHash
func VerifyPassword(hashedPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password));
	if err != nil {
		return false, errors.New("passwords do not match")
	}
	return true, nil
}
