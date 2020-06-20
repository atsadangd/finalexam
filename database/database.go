package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	createTb := `CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table customers", err)
	}

	fmt.Println("Create table success.")
}

func Conn() *sql.DB {
	return db
}
