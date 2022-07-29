package handler

import (
	"api-ranufrozen/database"
	"api-ranufrozen/food"
	"api-ranufrozen/merchant"
	"api-ranufrozen/order"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartRestAPI() {
	db := database.GetRDBConn()

	foodRepository := food.NewRepository(db)

	foodHandler := NewFoodHandler(foodRepository)

	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository)
	orderHandler := NewOrderHandler(orderService)

	merchantRepo := merchant.NewRepository(db)
	merchantHandler := NewMerchantHandler(merchantRepo)
	// 1 order punya 1 bill
	// bill: belum tau cara menggunakannya
	// billRepo := bill.NewRepository(db)
	// billService := bill.NewService(billRepo)
	// billHandler := handler.NewBillHandler(billService)

	// custRepository := customer.NewRepository(db)
	// custService := customer.NewService(db)
	// custHandler := customer.NewRepository(db)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "Ranufrozen",
			"panduan": "silahkan buka base_url/v1/url",
		})
	})

	v1 := router.Group("/v1")

	{
		v1.GET("/food/:id", foodHandler.Show)
		v1.GET("/foods/order_by", foodHandler.OrderBy)
		v1.POST("/order", orderHandler.PostorderHandler)

		v1.GET("/order/:id", orderHandler.Show)
		v1.GET("/orders/order_by", orderHandler.OrderBy)
		v1.POST("/order", orderHandler.PostorderHandler)

		v1.GET("/", merchantHandler.RootHandler)
		v1.GET("/merchant/:id", merchantHandler.Show)
		v1.GET("/merchants/order_by", merchantHandler.OrderBy)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/food/:id", foodHandler.Show)
	}
	fmt.Println("api running on port 8080")
	router.Run()
}
