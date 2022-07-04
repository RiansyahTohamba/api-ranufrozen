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
