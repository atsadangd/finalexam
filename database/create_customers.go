package database

import ("fmt")

func CreateCustomer(name, email, status string) (int, error) {
	var id int	
	row := Conn().QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3)  RETURNING id", name, email, status)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("can't create customers: %w", err)
	}

	return id, nil
}