package main

import (
	"api-ranufrozen/food"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connect error")
	}

	foodRepository := food.NewRepository(db)
	foodService := food.NewService(foodRepository)
	foodHandler := handler.NewFoodHandler(foodService)

	// foods, err := foodRepository.FindAll()

	// for _, food := range foods {
	// 	fmt.Println("food :", food.Name)
	// }

	// buat migrasi kode entity to table
	// db.AutoMigrate(&food.Food{})

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", foodHandler.RootHandler)
	v1.GET("/food/:id", foodHandler.Show)
	v1.GET("/foods/order_by", foodHandler.OrderBy)
	v1.POST("/food", foodHandler.PostFoodHandler)

	v2 := router.Group("/v2")
	// apakah handler nya dipisahkan saja?
	v2.GET("/food/:id", foodHandler.Show)

	router.Run()
}
