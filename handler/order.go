package handler

import (
	"api-ranufrozen/food"
	"api-ranufrozen/order"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type orderHandler struct {
	orderService order.Service
}

// function NeworderHandler bukan punya struct, tapi punya food.handler
// saat dipanggil jadi seperti ini `orderService := food.NewService(foodRepository)`
func NewOrderHandler(orderService order.Service) *orderHandler {
	return &orderHandler{orderService}
}

// func NewFoodHandler(foodService food.Service) *foodHandler {
// 	return &foodHandler{foodService}
// }

// v1.GET("/", orderHandler.RootHandler)
// v1.GET("/order/:id", orderHandler.Show)
// v1.GET("/orders/order_by", orderHandler.OrderBy)

// v1.POST("/order", orderHandler.PostorderHandler)

// public method diawali huruf CAPITAL
// RootHandler adalah public method
func (handler *orderHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Ranufrozen",
		"tagline": "makan enak untuk semua!",
	})
}

// Param untuk case base_url/:param
func (handler *orderHandler) Show(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// example case for 'query handler'
// base_url/foods?id=12
func (handler *orderHandler) OrderBy(c *gin.Context) {
	field := c.Query("field")

	foods, err := handler.orderService.FindAll()

	if err != nil {
		fmt.Println(err)
	}
	// for _, food := range foods {
	// 	fmt.Println("food :", food.Name)
	// }

	c.JSON(http.StatusOK, gin.H{
		"field": field,
		"foods": foods,
	})
}

// body
func (handler *orderHandler) PostorderHandler(c *gin.Context) {
	var foodReq food.FoodRequest

	err := c.ShouldBindJSON(&foodReq)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMsg)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}

	// food, err := handler.orderService.Create(foodReq)

	// jika terjadi error pada DB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		// "data": food,
	})

}
