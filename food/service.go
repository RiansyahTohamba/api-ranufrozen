package food

import (
	"fmt"
)

type Service interface {
	// FindById(id int) (Food, error)
	// FindAll() ([]Food, error)
	Create(food FoodRequest) (Food, error)
}

// struct vs interface?
// type
type service struct {
	foodRepo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) FindAll() ([]Food, error) {
	foods, err := s.foodRepo.FindAll()
	return foods, err
}

func (ser *service) OptimisTx() {
	prodId := 1
	quantity := 4
	fmt.Printf("Beli Produk sebanyak %d \n", quantity)
	ser.foodRepo.BuyProduct(prodId, quantity)

}

func (s *service) PrintFindAll() {
	foods, err := s.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, val := range foods {
		fmt.Println(val)
	}
}

func (ser *service) PrintProduct(id int) {
	food := ser.Get(id)
	fmt.Println(food)

	fmt.Println("food.Stock terbaru")
	fmt.Println(food.Stock)

}

// func (s *service) FindById(id int) (Food, error) {
// 	food, err := s.foodRepo.FindById(id)
// 	return food, err
// }

func (s *service) Create(foodReq FoodRequest) (Food, error) {
	price, _ := foodReq.Price.Float64()
	food := Food{
		Name:  foodReq.Name,
		Price: float64(price),
	}
	// disini harus ada pengondisian untuk price 'less or equal' dibawah zero
	newFood, err := s.foodRepo.Create(food)
	return newFood, err
}

func (ser service) Get(id int) Food {
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
