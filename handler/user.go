package handler

import (
	"api-ranufrozen/account"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

var secretKey = []byte("rahasia-loh")

type user struct {
	accountRepo *account.Repository
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserHandler(accountRepo *account.Repository) *user {
	return &user{accountRepo}
}

func (us *user) Show(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"username": ctx.GetString("username"),
		"data":     "lisf of data user",
	})
}

func (us *user) Login(ctx *gin.Context) {
	var usrReq = UserRequest{}
	ctx.ShouldBindJSON(&usrReq)
	username, err := us.accountRepo.Login(usrReq.Username, usrReq.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := generateJwt(username)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("parsing JWT")
		tknHeader := ctx.GetHeader("Authorization")
		if tknHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "no token provided",
			})
		}

		removeStr := (len("Bearer "))
		tknHeader = tknHeader[removeStr:]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tknHeader, claims, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			log.Println(err)
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token is invalid: renew your token",
			})
		}
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}

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
