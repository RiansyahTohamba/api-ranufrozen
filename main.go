package main

import (
	"api-ranufrozen/database"
	"api-ranufrozen/drink"
	"api-ranufrozen/food"
	"api-ranufrozen/handler"
	"fmt"
)

func main() {
	// cli()
	handler.StartRestAPI()
}
func cli() {
	rdb := database.GetRDBConn()
	rcl := database.GetRedisConn()
	foodRepository := food.NewRepository(rdb)
	foodCli := food.NewCli(foodRepository, rcl)
	foodCli.PrintProducts()
}

func mongoExample() {
	mongoCon := database.GetMongoConn()
	drinkRep := drink.NewDrinkRepo(mongoCon)
	drinkCli := drink.NewCli(*drinkRep)
	// 1. Create Drink, many Drink
	// InsetSampleDrink(db)

	// 2. Retrieve Specific Drink
	id := "62bd7b4ab1cf5abe26fb7e6b"
	fmt.Println(drinkCli.Show(id))

	// 3. Retrieve All Drink
	fmt.Println(drinkCli.List())
}
