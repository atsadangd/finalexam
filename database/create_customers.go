package database

import (
	"fmt"

	"github.com/atsadangd/finalexam/types"
)

func CreateCustomer(cus types.Customer) (int, error) {
	var id int
	row := Conn().QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3)  RETURNING id", cus.Name, cus.Email, cus.Status)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("can't create customers: %w", err)
	}

	return id, nil
}
