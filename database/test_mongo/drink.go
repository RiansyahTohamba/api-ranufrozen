package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type drinkRepository struct {
	db *mongo.Database
}

func NewDrinkRepo(db *mongo.Database) *drinkRepository {
	return &drinkRepository{db}
}

func (dr drinkRepository) aggregate(params string) (int, error) {
	// menghitung frekuensi tiap minuman per nama?
	return 0, nil
}
func (dr drinkRepository) paginate(offset, limit int) ([]Drink, error) {
	return []Drink{}, nil
}

func (dr drinkRepository) findById() ([]Drink, error) {
	// var ctx = context.Background()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	return nil, errors.New("error found")
	// }

	return Drink{}, nil
}

func (dr drinkRepository) findAll() ([]Drink, error) {

	var ctx = context.Background()

	csr, err := dr.db.Collection("drinks").Find(ctx, bson.D{{}})

	if err != nil {
		log.Fatal(err.Error())
		return nil, errors.New("error found")
	}

	defer csr.Close(ctx)

	var row Drink
	res := make([]Drink, 0)

	for csr.Next(ctx) {
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
			return nil, errors.New("error found")
		}
		res = append(res, row)
	}

	return res, nil
}

func (dr drinkRepository) insert(drink Drink) {
	var ctx = context.Background()
	_, err := dr.db.Collection("drinks").InsertOne(ctx, drink)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}

type Drink struct {
	Name string `bson:"name"`
}
