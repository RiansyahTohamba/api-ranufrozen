package main

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
)

// PARSE WITH CLAIMS

type Claims struct {
	Username string
	jwt.StandardClaims
}

var secretKey = []byte("rahasialoh")

func parseJwt(tokenInput string) string {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenInput, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Println(err)
		return "invalid token"
	}
	if !token.Valid {
		return "invalid token, check if expired or not"
	}
	return claims.Username
}

func main() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InJpYW4iLCJleHAiOjE2NTkzMjYwMzN9.GVR-d0YLsg_2jtfs7eZyIjl8mMRMTeXUHm82_yuMZnw"
	fmt.Println(parseJwt(token))
}
