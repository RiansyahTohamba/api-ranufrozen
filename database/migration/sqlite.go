package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := Migrate()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
}
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database/ranufrozen.db")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL			
		);
	`)
	if err != nil {
		return nil, err
	}
	return db, nil
}
