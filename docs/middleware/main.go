package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func m1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("keym1", "from m1")
	}
}
func m2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fromM1 := ctx.GetString("keym1")
		log.Println(fromM1)
	}
}

func main() {
	router := gin.Default()

	authorized := router.Use(m1()).Use(m2())
	authorized.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "hai",
		})
	})
	router.Run()
}
