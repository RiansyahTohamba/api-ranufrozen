package handler

import (
	"api-ranufrozen/food"
	"log"
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

// kalau c.Query() untuk apa?
func FoodHandler(c *gin.Context) {
	name := c.Query("name")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"price": price,
	})
}

// example query handler
// base_url/foods?id=12
func Show(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func PostFoodHandler(c *gin.Context) {
	var foodInput food.FoodInput
	err := c.ShouldBindJSON(&foodInput)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"name":       foodInput.Name,
		"photo_path": foodInput.PhotoPath,
		"price":      foodInput.Price,
	})
}
