package user

import "github.com/gin-gonic/gin"

type user struct {
}

func NewUserHandler() *user {
	return &user{}
}

func (us *user) Show(ctx gin.Context) {

}
