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

func (cl *cli) Sum(num1, num2 int) int {
	return 0
}

func (cl *cli) OptimisTx() {
	prodId := 1
	quantity := 4
	fmt.Printf("Beli Produk sebanyak %d \n", quantity)
	cl.foodRepo.BuyProduct(prodId, quantity)

}

func (cl *cli) findAll() []Food {
	// get foods from redis
	ctx := context.Background()
	foodJson, err := cl.rcl.Get(ctx, "foods").Bytes()

	// food json tidak terdapat pada cache
	if err != nil {
		offset := 5
		limit := 8
		foodArr, err := cl.foodRepo.FindAll(offset, limit)

		if err != nil {
			log.Println(err)
		}

		cl.setCache(ctx, foodArr)

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

func (cl *cli) setCache(ctx context.Context, foodArr []Food) {
	foodJson, err := json.Marshal(foodArr)

	if err != nil {
		log.Println(err)
	}

	expired := 10 * time.Second
	err = cl.rcl.Set(ctx, "foods", foodJson, expired).Err()

	if err != nil {
		log.Println(err)
	}
}

func (cl *cli) PrintFoods() {
	foods := cl.findAll()
	for _, food := range foods {
		fmt.Println(food.ID, food.Name)
	}
}

func (cl *cli) PrintFood(id int) {
	food, err := cl.foodRepo.FindById(id)

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(food.ID, food.Name)
	fmt.Println("Stock terkini: ")
	fmt.Println(food.Stock)
}

func (cl *cli) Create(foodReq FoodRequest) (Food, error) {
	price, _ := foodReq.Price.Float64()
	food := Food{
		Name:  foodReq.Name,
		Price: float64(price),
	}
	// disini harus ada pengondisian untuk price 'less or equal' dibawah zero
	newFood, err := cl.foodRepo.Create(food)
	return newFood, err
}

// func GetMock(id int) Food {
// 	food := Food{
// 		name:  "Nugget",
// 		price: 20000,
// 	}
// 	return food

// }
