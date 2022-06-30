package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db, err := connect()
	if err != nil {
		log.Println(err)
	}

	// ==== food example ======
	// foodRep := NewFoodRepo(db)
	// foodRep.find()

	// ==== Drink example ======
	drinkRep := NewDrinkRepo(db)

	// 1. Create Drink, many Drink
	// InsetSampleDrink(db)

	// 2. Retrieve Specific Drink
	fmt.Println(drinkRep.find())

	// 3. Retrieve All Drink
	fmt.Println(drinkRep.findAll())

	// 4. Update Drink
	// 4.a Retrieve Specific Drink
	// 5. Delete Drink 1

}

func InsetSampleDrink(db *mongo.Database) {
	drinkRep := NewDrinkRepo(db)

	drinkRep.insert(Drink{"Coca-cola"})
	drinkRep.insert(Drink{"Fanta"})
	drinkRep.insert(Drink{"Teh Botol"})
	drinkRep.insert(Drink{"Teh Kotak"})
}
func connect() (*mongo.Database, error) {
	var ctx = context.Background()

	clOpt := options.Client()
	clOpt.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clOpt)
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}
	return client.Database("ranufrozen"), nil
}
