package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type foodRepository struct {
	db *mongo.Database
}

func NewFoodRepo(db *mongo.Database) *foodRepository {
	return &foodRepository{db}
}

func (fr foodRepository) find() {
	var ctx = context.Background()

	csr, err := fr.db.Collection("foods").Find(ctx, bson.D{{}})

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

// relational
