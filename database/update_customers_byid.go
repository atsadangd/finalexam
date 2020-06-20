package database

import (
	"fmt"

	"github.com/atsadangd/finalexam/types"
)

func UpdateCustomersById(cus types.Customer) error {
	stmt, err := Conn().Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1;")
	if err != nil {
		return fmt.Errorf("can't prepare update statement: %w", err)
	}

	if _, err := stmt.Exec(cus.ID, cus.Name, cus.Email, cus.Status); err != nil {
		return fmt.Errorf("can't execute update statement: %w", err)
	}

	return nil
}
