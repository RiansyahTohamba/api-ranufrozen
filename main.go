package main

import (
	"api-ranufrozen/handler"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	errDotenv := godotenv.Load()
	if errDotenv != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connect error")
	}

	// db.AutoMigrate(&food.Food{})

	router := gin.Default()

	v1 := router.Group("/v1")

	// var food food.Food
	// err = db.Debug().where
	// router := gin.Default()
	v1.GET("/", handler.RootHandler)
	v1.GET("/food/:id", handler.Show)
	v1.GET("/foods/order_by", handler.OrderBy)
	v1.POST("/food", handler.PostFoodHandler)

	router.Run()
}
