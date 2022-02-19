package handler

import (
	"api-ranufrozen/order"
)

type orderHandler struct {
	orderService order.Service
}

// function NewFoodHandler bukan punya struct, tapi punya food.handler
// saat dipanggil jadi seperti ini `foodService := food.NewService(foodRepository)`
func NewOrderHandler(orderService order.Service) *orderHandler {
	return &orderHandler{orderService}
}
