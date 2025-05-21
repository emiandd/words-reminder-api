package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var jwtKey = []byte("access_token")

func CreateToken(id int64, email string) (string, error) {
	claims := &Claims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Tiempo de expiración del token
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

