package database

import "fmt"

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
