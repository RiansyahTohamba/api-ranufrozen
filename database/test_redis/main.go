package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

func main() {
	ctx := context.Background()
	// run on ubuntu redis-server --daemonize yes
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

}
