package main

import (
	"context"
	"fmt"
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
