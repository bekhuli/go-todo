package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return CreateUser(username, string(hashedPassword))
}
