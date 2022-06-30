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

var ctx = context.Background()

func connect() (*mongo.Database, error) {
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

func find() {
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

func main() {
	find()
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
