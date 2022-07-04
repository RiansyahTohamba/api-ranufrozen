package main

import (
	"context"
	"fmt"
)

var ctx = context.TODO()

func main() {
	rdsClient := GetRedisConn()
	// cartExample(rdsClient)
	stringExample(rdsClient)

}

func stringExample(rdsClient *RedisClient) {
	// SET
	// bagaimana jika SET 2 kali?
	stRep := NewStringRepo(rdsClient)
	insertStrRep(stRep)
	// GET
	// data terakhir yang akan diambil
	fmt.Println(stRep.Show("user1"))
	fmt.Println(stRep.Show("user2"))
	fmt.Println(stRep.Show("user3"))
	// lalu jika ada key yang sama pada fitur cart sebelumnya (cart menggunakan key 'user1')
	// maka akan ditimpa dengan fitur sesudahnya

}

func hashExample(rdsClient *RedisClient) {
}

func cartExample(rdsClient *RedisClient) {
	cartRepo := NewCart(rdsClient)

	insertCart(*cartRepo)
	// GET
	fmt.Println(cartRepo.GetCart(ctx, "user1"))
	fmt.Println(cartRepo.GetCart(ctx, "user2"))

	removeCart(*cartRepo)
	fmt.Println("after deletion")
	fmt.Println(cartRepo.GetCart(ctx, "user1"))
	fmt.Println(cartRepo.GetCart(ctx, "user2"))

}

func insertCart(cartRepo CartRepository) {
	// SET
	// User1 Buy Food
	cartRepo.AddCart(ctx, "user1", "buah naga")
	// buah naga nya akan tertimpa dengan nasi kucing
	cartRepo.AddCart(ctx, "user1", "nasi kucing")

	// User2 Buy Drink
	cartRepo.AddCart(ctx, "user2", "latte coffee")
	cartRepo.AddCart(ctx, "user2", "latte tea")
}

func insertStrRep(stRep *StringRepo) {
	stRep.Create("user1", "rendang")
	stRep.Create("user1", "mie goreng")
	stRep.Create("user2", "ayam bakar")
	stRep.Create("user3", "roti panggang")
}

func removeCart(cartRepo CartRepository) {
	cartRepo.DeleteCart(ctx, "user1")
	cartRepo.DeleteCart(ctx, "user2")
}
