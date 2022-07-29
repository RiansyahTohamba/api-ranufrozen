package food

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

// struct vs interface?
// type
type cli struct {
	foodRepo Repository
	rcl      *redis.Client
}

func NewCli(repo Repository, rcl *redis.Client) *cli {
	return &cli{repo, rcl}
}

func (s *cli) Sum(num1, num2 int) int {
	return 0
}

func (ser *cli) OptimisTx() {
	prodId := 1
	quantity := 4
	fmt.Printf("Beli Produk sebanyak %d \n", quantity)
	ser.foodRepo.BuyProduct(prodId, quantity)

}

func (s *cli) foodsJson() []Food {
	// get foods from redis
	ctx := context.Background()
	foodJson, err := s.rcl.Get(ctx, "foods").Bytes()

	// food json tidak terdapat pada cache
	if err != nil {
		foodArr, err := s.foodRepo.FindAll()

		if err != nil {
			log.Println(err)
		}

		s.setCache(ctx, foodArr)

		log.Println("from SQL")

		return foodArr
	}

	foodArr := []Food{}
	// mapping foods of json to array of foods
	err = json.Unmarshal(foodJson, &foodArr)

	if err != nil {
		log.Println(err)
	}

	log.Println("from redis")
	return foodArr

}

func (s *cli) setCache(ctx context.Context, foodArr []Food) {
	foodJson, err := json.Marshal(foodArr)

	if err != nil {
		log.Println(err)
	}

	expired := 10 * time.Second
	err = s.rcl.Set(ctx, "foods", foodJson, expired).Err()

	if err != nil {
		log.Println(err)
	}
}

func (s *cli) PrintProducts() {
	foods := s.foodsJson()
	for _, food := range foods {
		fmt.Println(food.ID, food.Name)
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
