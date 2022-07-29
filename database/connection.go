package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// key-value
func GetRedisConn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}
func GetMongoConn() *mongo.Database {
	var ctx = context.Background()

	clOpt := options.Client()
	clOpt.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clOpt)
	if err != nil {
		log.Println(err)
		return nil
	}
	err = client.Connect(ctx)

	if err != nil {
		log.Println(err)
		return nil
	}
	return client.Database("ranufrozen")
}

func GetRDBConn() *gorm.DB {
	errDotenv := godotenv.Load()
	if errDotenv != nil {
		log.Fatal("Error loading .env file")
	}

	appEnv := os.Getenv("APP_ENV")

	var db *gorm.DB
	var err error

	if appEnv == "development" {
		db, err = getSqliteConn()
	} else if appEnv == "production" {
		db, err = getMysqlConn()
	}

	if err != nil {
		log.Fatal("DB connect error")
	}
	return db
}

func getSqliteConn() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database/ranufrozen.db"), &gorm.Config{})
	return db, err
}

func getMysqlConn() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("running : sudo systemctl start mysql")
		log.Println(err)
	}
	return db, nil
}
