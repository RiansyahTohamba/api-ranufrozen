package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	serverMongo   = "172.17.0.2"
	userMysql     = "postgres"
	passwordMysql = "postgres"
	mongoDB       = "jug"
	mysqlDB       = "jug"
	sqlitePath    = ""
)

// key-value
func GetRedisConn() {

}
func GetMongoConn() {

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
	return db, err
}
