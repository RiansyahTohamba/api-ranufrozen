package main

import (
	"api-ranufrozen/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/api-ranufrozen?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connect error")
	}
	fmt.Println("DB connection succed")

	router := gin.Default()

	v1 := router.Group("/v1")

	// var food food.Food
	// err = db.Debug().where
	// router := gin.Default()
	v1.GET("/", handler.RootHandler)

	router.Run()
}
