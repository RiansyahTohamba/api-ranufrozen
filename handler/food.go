package handler

import (
	"api-ranufrozen/food"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type foodHandler struct {
	foodService food.Service
}

// function NewFoodHandler bukan punya struct, tapi punya food.handler
// saat dipanggil jadi seperti ini `foodService := food.NewService(foodRepository)`
func NewFoodHandler(foodService food.Service) *foodHandler {
	return &foodHandler{foodService}
}

// func (h *foodHandler) GetFoods(c *gin.Context) {
// 	foods, err := h.foodService.FindAll()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err,
// 		})
// 		return
// 	}
// 	var foodsResponse []food.FoodResponse
// 	for _, b := range foods {
// 		foodResponse := convertToBookResponse(b)
// 		foodsResponse = append(foodsResponse, foodResponse)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": foodsResponse,
// 	})
// }

// func (h *foodHandler) CreateFood(c *gin.Context) {
// 	var foodRequest food.FoodRequest

// }

// public method diawali huruf CAPITAL
// RootHandler adalah public method
func (handler *foodHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Ranufrozen",
		"tagline": "makan enak untuk semua!",
	})
}

// Param untuk case base_url/:param
func (handler *foodHandler) Show(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// example case for 'query handler'
// base_url/foods?id=12
// func (handler *foodHandler) OrderBy(c *gin.Context) {
// 	field := c.Query("field")

// 	foods, err := handler.foodService.FindAll()

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// for _, food := range foods {
// 	// 	fmt.Println("food :", food.Name)
// 	// }

// 	c.JSON(http.StatusOK, gin.H{
// 		"field": field,
// 		"foods": foods,
// 	})
// }

// body
func (handler *foodHandler) PostFoodHandler(c *gin.Context) {
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
	food, err := handler.foodService.Create(foodReq)

	// jika terjadi error pada DB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": food,
	})

}
