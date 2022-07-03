package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis/v9"
)

func main() {
	ctx := context.Background()

	kvdb := GetRedisConn()
	err := kvdb.Set(ctx, "key", "value", 0).Err()

	if err != nil {
		fmt.Println("running redis-server --daemonize yes")
		panic(err)
	}

	val, err := kvdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

}

func GetRedisConn() *redis.Client {
	var once sync.Once
	var kvdb *redis.Client
	var ctx = context.Background()

	once.Do(func() {
		kvdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	})

	_, err := kvdb.Ping(ctx).Result()

	if err != nil {
		fmt.Println("running redis-server --daemonize yes")
		log.Fatal(err)
	}

	return kvdb
}
