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

/*
func CreateCustomer(name, email, status string) error {
	row := Conn().QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3)  RETURNING id", name, email, status)
	cus := Customer{}
	err := row.Scan(&cus.ID)
	if err != nil {
		return fmt.Errorf("can't create customers: %w", err)
	}

	return nil
}
*/

func DeleteCustomersById(id string) error {
	stmt, err := Conn().Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		return fmt.Errorf("can't prepare delete statement: %w", err)
	}

	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("can't execute delete statement: %w", err)
	}

	return nil
}
