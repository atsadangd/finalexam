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

func CreateCustomer(name, email, status string) (error, int) {
	var id int
	row := Conn().QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3)  RETURNING id", name, email, status)
	err := row.Scan(&id)
	if err != nil {
		return fmt.Errorf("can't create customers: %w", err), 0
	}

	return nil, id
}

func GetCustomersByID(id string) (error, *sql.Row) {
	stmt, err := Conn().Prepare("SELECT id, name, email, status FROM customers where id=$1")
	if err != nil {
		return fmt.Errorf("can't prepare select statement: %w", err), nil
	}

	row := stmt.QueryRow(id)
	return nil, row
}

func GetCustomers() (error, *sql.Rows) {
	stmt, err := Conn().Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		return fmt.Errorf("can't prepare select statement: %w", err), nil
	}

	rows, err := stmt.Query()
	if err != nil {
		return fmt.Errorf("can't get customers data: %w", err), nil
	}

	return nil, rows
}

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

func UpdateCustomersById(id, name, email, status string) error {
	stmt, err := Conn().Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1;")
	if err != nil {
		return fmt.Errorf("can't prepare update statement: %w", err)
	}

	if _, err := stmt.Exec(id, name, email, status); err != nil {
		return fmt.Errorf("can't execute update statement: %w", err)
	}

	return nil
}
