package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err = Migrate()
	if err != nil {

	}
}
func Migrate() {
	db, err := sql.Open("sqlite3", "database")
	if err != nil {
		log.Println(err)
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

	}
}
