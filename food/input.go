package food

type FoodInput struct {
	// json key nya 'name'
	Name string `json:"name" binding:"required" `
	// PhotoPath string
	// Rating    int
	Price float64 `json:"price" binding:"required,number" `
}
