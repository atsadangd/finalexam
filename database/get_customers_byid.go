package database

import (
	"database/sql"
	"fmt"
)

func GetCustomersByID(id string) (*sql.Row, error) {
	stmt, err := Conn().Prepare("SELECT id, name, email, status FROM customers where id=$1")
	if err != nil {
		return nil, fmt.Errorf("can't prepare select statement: %w", err)
	}

	row := stmt.QueryRow(id)
	return row, nil
}
