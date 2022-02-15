package food

type FoodInput struct {
	Name      string `json:"name" binding:"" `
	PhotoPath string
	Rating    int
	Price     float64
}
