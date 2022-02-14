package handler

import (
	"api-ranufrozen/food"
	"net/http"

	"github.com/gin-gonic/gin"
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

func HelloWorld(name string) string {
	return ("Hello " + name)
}

// func FoodHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"name":    "Ranufrozen",
// 		"tagline": "makan enak untuk semua!",
// 	})
// }

func PostFoodHandler(c *gin.Context) {
	// var foodRepo food.Repository
	// Create

}
