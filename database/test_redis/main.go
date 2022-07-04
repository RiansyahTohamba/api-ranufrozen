package main

import (
	"context"
	"fmt"
)

var ctx = context.TODO()

func main() {
	rClient := GetRedisConn()
	insertData(*rClient)
	// GET
	fmt.Println(rClient.GetCart(ctx, "user1"))
	fmt.Println(rClient.GetCart(ctx, "user2"))

	removeData(*rClient)
	fmt.Println("after deletion")
	fmt.Println(rClient.GetCart(ctx, "user1"))
	fmt.Println(rClient.GetCart(ctx, "user2"))

}

func insertData(rClient RedisClient) {
	// SET
	// User1 Buy Food
	rClient.AddCart(ctx, "user1", "buah naga")
	rClient.AddCart(ctx, "user1", "nasi kucing")

	// User2 Buy Drink
	rClient.AddCart(ctx, "user2", "latte coffee")
	rClient.AddCart(ctx, "user2", "latte tea")
}

func removeData(rClient RedisClient) {
	rClient.DeleteCart(ctx, "user1")
	rClient.DeleteCart(ctx, "user2")
}
