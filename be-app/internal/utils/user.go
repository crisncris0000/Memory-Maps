package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error hasing password", err)
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func GenerateJWTToken(email string, role int) *jwt.Token {
	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24)

	return token
}
