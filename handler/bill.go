package handler

import (
	"api-ranufrozen/bill"
)

type billHandler struct {
	billRepo bill.Repository
}

func NewBillHandler(billRepo bill.Repository) *billHandler {
	return &billHandler{billRepo}

}
