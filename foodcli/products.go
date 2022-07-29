package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
	_ "github.com/mattn/go-sqlite3"
)

type Products struct {
	ProductId   int     `json:"id"`
	ProductName string  `json:"name"`
	RetailPrice float64 `json:"price"`
}

type JsonResponse struct {
	Data   []Products `json:"data"`
	Source string     `json:"source"`
}

func getProducts() (*JsonResponse, error) {
	// redis tidak cocok disimpan di repository
	// karena redis sudah menyimpan dalam bentuk JSON response
	// repository hanya return array of data
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	cachedProducts, err := redisClient.Get(ctx, "products").Bytes()

	response := JsonResponse{}

	if err != nil {

		dbProducts, err := fetchFromDb()

		if err != nil {
			return nil, err
		}

		cachedProducts, err = json.Marshal(dbProducts)

		if err != nil {
			return nil, err
		}
		expiredData := 30 * time.Second
		err = redisClient.Set(ctx, "products", cachedProducts, expiredData).Err()

		if err != nil {
			return nil, err
		}

		response = JsonResponse{Data: dbProducts, Source: "Disk SQL"}

		return &response, err
	}

	products := []Products{}

	err = json.Unmarshal(cachedProducts, &products)
	// staleness
	if err != nil {
		return nil, err
	}

	response = JsonResponse{Data: products, Source: "Redis Cache"}

	return &response, nil
}

func fetchFromDb() ([]Products, error) {
	db, err := sql.Open("sqlite3", "ranufrozen.db")
	if err != nil {
		return nil, err
	}

	queryString := `select id,name, price from foods`

	rows, err := db.Query(queryString)

	if err != nil {
		return nil, err
	}

	var records []Products

	for rows.Next() {

		var p Products

		err = rows.Scan(&p.ProductId, &p.ProductName, &p.RetailPrice)

		records = append(records, p)

		if err != nil {
			return nil, err
		}

	}

	return records, nil
}
