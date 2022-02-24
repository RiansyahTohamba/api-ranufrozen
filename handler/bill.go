package handler

import (
	"api-ranufrozen/bill"
)

type billHandler struct {
	billService bill.Service
}

// function NewFoodHandler bukan punya struct, tapi punya food.handler
// saat dipanggil jadi seperti ini `foodService := food.NewService(foodRepository)`
func NewBillHandler(billService bill.Service) *billHandler {
	return &billHandler{billService}

}
