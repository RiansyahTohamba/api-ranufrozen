package main

import (
	"api-ranufrozen/database"
	"api-ranufrozen/drink"
	"api-ranufrozen/food"
	"api-ranufrozen/handler"
	"api-ranufrozen/order"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cli()
	// restAPI()
}

func exampleCase() {

}

func cli() {
	// foodRdb()
	// drinkDocBased()
	shopCartKv()
}

func shopCartKv() {
	// ==== shopping cart
	// connect redis terlebih dahulu
	kvdb := database.GetRedisConn()
	fmt.Println(kvdb)
	key := "key-1"
	data := "Hello Redis"
	ttl := time.Duration(3) * time.Second
	// store data using SET command
	op1 := kvdb.Set(context.Background(), key, data, ttl)

	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("set operation success")

	// Buy Food
	// Buy Drink
	// pertanyaan terkait:
	// 1. disimpan dimana feature nya?
	// 2. apakah perlu repo?
}
func drinkDocBased() {
	// ==== Drink example ======
	// 1. Create Drink, many Drink
	// InsetSampleDrink(db)
	// 2. Retrieve Specific Drink
	// 3. Retrieve All Drink
	mongoCon := database.GetMongoConn()

	drinkRep := drink.NewDrinkRepo(mongoCon)
	drinkCli := drink.NewCli(*drinkRep)

	id := "62bd7b4ab1cf5abe26fb7e6b"
	fmt.Println(drinkCli.Show(id))

	fmt.Println(drinkCli.List())
}

func foodRdb() {
	rdb := database.GetRDBConn()
	foodRepository := food.NewRepository(rdb)
	foodCli := food.NewCli(foodRepository)

	foodCli.PrintFindAll()
	// foodCli.OptimisTx()
	// foodCli.PrintProduct(1)
}

func restAPI() {
	// DB_PASSWORD = "AAAA"
	// sudo systemctl start mysql
	db := database.GetRDBConn()

	foodRepository := food.NewRepository(db)
	// food = new Food()
	// food.
	// foodService := food.NewService(foodRepository)
	foodHandler := handler.NewFoodHandler(foodRepository)

	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	// merchantRepo := merchant.NewRepository(db)
	// merchantServ := merchant.NewService(merchantRepo)
	// merchantHandler := handler.NewMerchantHandler(merchantServ)
	// 1 order punya 1 bill
	// bill: belum tau cara menggunakannya
	// billRepo := bill.NewRepository(db)
	// billService := bill.NewService(billRepo)
	// billHandler := handler.NewBillHandler(billService)

	// custRepository := customer.NewRepository(db)
	// custService := customer.NewService(db)
	// custHandler := customer.NewRepository(db)

	// buat migrasi kode entity to table
	// db.AutoMigrate(&food.Food{})

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "Ranufrozen",
			"panduan": "silahkan buka base_url/v1/url",
		})
	})

	v1 := router.Group("/v1")

	v1.GET("/food/:id", foodHandler.Show)
	// v1.GET("/foods/order_by", foodHandler.OrderBy)
	// v1.POST("/order", orderHandler.PostorderHandler)

	// fitur order, halaman apa saja yang mungkin muncul?
	// 1. as a pelanggan?
	// apakah sudah ada template html nya?
	//
	// 2. as A Customer Service (CS)
	// ini untuk dashboard CS nantinya
	// siapa tau ada masalah dengan transaksi
	// jadi CS Punya data

	v1.GET("/order/:id", orderHandler.Show)
	v1.GET("/orders/order_by", orderHandler.OrderBy)
	// v1.POST("/order", orderHandler.PostorderHandler)

	// v1.GET("/", merchantHandler.RootHandler)
	// v1.GET("/merchant/:id", merchantHandler.Show)
	// v1.GET("/merchants/order_by", merchantHandler.OrderBy)

	// ========================= V2 =====================
	v2 := router.Group("/v2")
	// apakah handler nya dipisahkan saja?
	v2.GET("/food/:id", foodHandler.Show)
	fmt.Println("api running on port 8080")
	router.Run()
}
