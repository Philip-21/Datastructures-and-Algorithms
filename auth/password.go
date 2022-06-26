package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(HashedPassword), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, HashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(password))
}
