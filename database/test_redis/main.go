package main

import (
	"context"
	"fmt"
)

var ctx = context.TODO()

func main() {
	rdsClient := GetRedisConn()
	cartRepo := NewCart(rdsClient)

	insertData(*cartRepo)
	// GET
	fmt.Println(cartRepo.GetCart(ctx, "user1"))
	fmt.Println(cartRepo.GetCart(ctx, "user2"))

	removeData(*cartRepo)
	fmt.Println("after deletion")
	fmt.Println(cartRepo.GetCart(ctx, "user1"))
	fmt.Println(cartRepo.GetCart(ctx, "user2"))

}

func insertData(cartRepo CartRepository) {
	// SET
	// User1 Buy Food
	cartRepo.AddCart(ctx, "user1", "buah naga")
	// buah naga nya akan tertimpa dengan nasi kucing
	cartRepo.AddCart(ctx, "user1", "nasi kucing")

	// User2 Buy Drink
	cartRepo.AddCart(ctx, "user2", "latte coffee")
	cartRepo.AddCart(ctx, "user2", "latte tea")
}

func removeData(cartRepo CartRepository) {
	cartRepo.DeleteCart(ctx, "user1")
	cartRepo.DeleteCart(ctx, "user2")
}
