// example_optimis.go
package main

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepo() {

}

type product struct {
	stock int
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	// sudo systemctl start mysql
	dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func (tr TransactionRepository) oldValue(id string) error {
	statement := `SELECT id, val1, val2 FROM theTable WHERE iD = ?`
	row, err = tr.db.QueryRow(statement, id)

	// - {code that calculates new values}
	newVal1 := 10 + 10

	// oldVal1 := 10

	statupdate := `UPDATE theTable 
		SET val1 = @newVal1
		WHERE iD = @theId AND 
		val1 = @oldVal1;`

	tx, err := tr.db.Begin()

	_, err = tx.Exec(statupdate, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if AffectedRows == 1 {
		//    {go on with your other code}
		fmt.Println("Perubahan success")
	} else {
		tx.Rollback()
		//    {decide what to do since it has gone bad... in your code}
	}
	return nil
}
