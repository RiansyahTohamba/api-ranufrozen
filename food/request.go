package food

import "encoding/json"

type FoodRequest struct {
	// json key nya 'name'
	Name string `json:"name" binding:"required" `
	// PhotoPath string
	// Rating    int
	Price json.Number `json:"price" binding:"required,number" `
}
