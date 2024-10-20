package utils

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)
import "github.com/golang-jwt/jwt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	signedString, err := token.SignedString([]byte("secret"))
	return "Bearer " + signedString, err
}
