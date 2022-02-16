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

// func NewFoodHandler(foodService food.Service) *foodHandler {
// 	return &foodHandler{foodService}
// }

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
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Ranufrozen",
		"tagline": "makan enak untuk semua!",
	})
}

// Param untuk case base_url/:param
func Show(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// example case for 'query handler'
// base_url/foods?id=12
func OrderBy(c *gin.Context) {
	field := c.Query("field")

	c.JSON(http.StatusOK, gin.H{
		"field": field,
	})
}

// body
func PostFoodHandler(c *gin.Context) {
	var foodInput food.FoodInput

	err := c.ShouldBindJSON(&foodInput)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMsg)
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"name":  foodInput.Name,
		"price": foodInput.Price,
	})
}
