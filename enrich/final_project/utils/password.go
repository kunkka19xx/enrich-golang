package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("Could not hash password %w", err)
	}

	return string(hashedPassword), nil
}

func verifyPassword(hashedPassword string, candicatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candicatePassword))
}
