package main

import (
	"context"
	"log"
)

type CartRepository struct {
	redisCl *RedisClient
}

func NewCart(redisCl *RedisClient) *CartRepository {
	return &CartRepository{redisCl}
}

// ini masih pakai data string, bagiamana dengan tipe data yang lain?
// harus baca referensi dulu utk mengetahui best practice membuat cart.
// fokus di main.py

// == CART function
func (cre *CartRepository) AddCart(ctx context.Context, cartId string, item string) {
	err := cre.redisCl.Set(ctx, cartId, item, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func (cre *CartRepository) GetCart(ctx context.Context, cartId string) string {
	val, err := cre.redisCl.Get(ctx, cartId).Result()
	if err != nil {
		log.Println(err)
	}
	return val
}

func (cre *CartRepository) DeleteCart(ctx context.Context, cartId string) {
	err := cre.redisCl.Del(ctx, cartId).Err()
	if err != nil {
		log.Println(err)
	}
}
