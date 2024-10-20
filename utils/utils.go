package utils

import (
	"errors"
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

// CheckPassword  验证密码
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	// err==nil 表示密码正确
	return err == nil
}

func ParseJwt(tokenString string) (string, error) {
	//如果有 去掉 Bearer 前缀
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("username claim  is not a string")
		}
		return username, nil
	}
	return "", err
}
