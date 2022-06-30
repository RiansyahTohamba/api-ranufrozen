package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 1. Create Drink
	// 1.a. Create Drink 2,3,4
	// 2. Retrieve Specific Drink
	// 3. Retrieve All Drink
	// 4. Update Drink
	// 4.a Retrieve Specific Drink
	// 5. Delete Drink 1
	db, err := connect()
	if err != nil {
		log.Println(err)
	}

	// food example
	foodRep := NewFoodRepo(db)
	foodRep.find()
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

type drinkRepository struct {
	db *mongo.Database
}

func NewDrinkRepo(db *mongo.Database) *drinkRepository {
	return &drinkRepository{db}
}

func (dr drinkRepository) find() {

}
func (dr drinkRepository) insert() {
	// var ctx = context.Background()

	// db, err := connect()

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// _, err = db.Collection("student").InsertOne(ctx, student{"Wick", 2})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// _, err = db.Collection("student").InsertOne(ctx, student{"Ethan", 2})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	fmt.Println("Insert success!")
}

type foodRepository struct {
	db *mongo.Database
}

func NewFoodRepo(db *mongo.Database) *foodRepository {
	return &foodRepository{db}
}

func (fr foodRepository) find() {
	var ctx = context.Background()

	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("foods").Find(ctx, bson.D{{}})

	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]Food, 0)

	for csr.Next(ctx) {
		var row Food
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}
	fmt.Println(result)

}

type Food struct {
	Name          string  `bson:"name"`
	PhotoPath     string  `bson:"photoPath"`
	Rating        int     `bson:"rating"`
	Price         float64 `bson:"price"`
	Stock         int     `bson:"stock"`
	IsSuperSeller int8    `bson:"isSuperSeller"`
	Category      int     `bson:"category"`
	QuantitySold  int     `bson:"qtSold"`
	Description   string  `bson:"desc"`
	Discount      float32

	CreatedAt time.Time
	UpdateAt  time.Time
}

type Drink struct {
	Name string `bson:"name"`
}
