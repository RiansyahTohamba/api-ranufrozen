package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// NEW WITH CLAIMS
// SIGNED STRING

type Claims struct {
	Username string
	jwt.StandardClaims
}

var secretKey = []byte("rahasialoh")

func generateJwt(username string) (string, error) {
	expiredTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secretKey)
	return tokenStr, err
}

func main() {
	username := "rian"
	token, _ := generateJwt(username)
	fmt.Println(token)
}
