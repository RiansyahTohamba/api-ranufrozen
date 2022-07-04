package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis/v9"
)

type RedisClient struct{ *redis.Client }

func GetRedisConn() *RedisClient {
	var once sync.Once
	var redisClient *RedisClient
	var ctx = context.Background()

	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		redisClient = &RedisClient{client}
	})

	_, err := redisClient.Ping(ctx).Result()

	if err != nil {
		fmt.Println("running redis-server --daemonize yes")
		log.Fatalf("Could not connect to redis %v", err)
	}
	return redisClient
}

// ini masih pakai data string, bagiamana dengan tipe data yang lain?
// harus baca referensi dulu utk mengetahui best practice membuat cart.
// fokus di main.py

// == CART function
func (rc *RedisClient) AddCart(ctx context.Context, cartId string, item string) {
	err := rc.Set(ctx, cartId, item, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func (rc *RedisClient) GetCart(ctx context.Context, cartId string) string {
	val, err := rc.Get(ctx, cartId).Result()
	if err != nil {
		log.Println(err)
	}
	return val
}
