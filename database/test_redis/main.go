package main

import (
	"context"
	"fmt"
)

var ctx = context.TODO()

func main() {
	rClient := GetRedisConn()
	// insertData(*rClient)
	// GET
	cartUser1 := rClient.GetCart(ctx, "user1")
	cartUser2 := rClient.GetCart(ctx, "user2")

	fmt.Println(cartUser1)
	fmt.Println(cartUser2)

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
