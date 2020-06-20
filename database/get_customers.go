package database

import (
	"database/sql"
	"fmt"
)

func GetCustomers() (*sql.Rows, error) {
	stmt, err := Conn().Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		return nil, fmt.Errorf("can't prepare select statement: %w", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("can't get customers data: %w", err)
	}

	return rows, nil
}
