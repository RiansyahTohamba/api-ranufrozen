package food

import (
	"fmt"
)

// struct vs interface?
// type
type cli struct {
	foodRepo Repository
}

func NewCli(repo Repository) *cli {
	return &cli{repo}
}

func (s *cli) Sum(num1, num2 int) int {
	return 0
}

func (s *cli) FindAll() ([]Food, error) {
	foods, err := s.foodRepo.FindAll()
	return foods, err
}

func (ser *cli) OptimisTx() {
	prodId := 1
	quantity := 4
	fmt.Printf("Beli Produk sebanyak %d \n", quantity)
	ser.foodRepo.BuyProduct(prodId, quantity)

}

func (s *cli) PrintFindAll() {
	foods, err := s.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, val := range foods {
		fmt.Println(val)
	}
}

func (ser *cli) PrintProduct(id int) {
	food := ser.Get(id)
	fmt.Println(food)

	fmt.Println("food.Stock terbaru")
	fmt.Println(food.Stock)

}

// func (s *cli) FindById(id int) (Food, error) {
// 	food, err := s.foodRepo.FindById(id)
// 	return food, err
// }

func (s *cli) Create(foodReq FoodRequest) (Food, error) {
	price, _ := foodReq.Price.Float64()
	food := Food{
		Name:  foodReq.Name,
		Price: float64(price),
	}
	// disini harus ada pengondisian untuk price 'less or equal' dibawah zero
	newFood, err := s.foodRepo.Create(food)
	return newFood, err
}

func (ser cli) Get(id int) Food {
	food := ser.foodRepo.FindById(id)
	return food

}

// func GetMock(id int) Food {
// 	food := Food{
// 		name:  "Nugget",
// 		price: 20000,
// 	}
// 	return food

// }
