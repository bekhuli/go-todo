package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/bekhuli/go-todo/config"

	"golang.org/x/crypto/bcrypt"
)

func Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return CreateUser(username, string(hashedPassword))
}

func Login(username, password string) (string, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	cfg := config.Envs
	expiration := time.Second * time.Duration(cfg.JWTExpirationInSeconds)
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}
