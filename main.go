package main

import "github.com/atsadangd/finalexam/customers"

func main() {
	r := customers.SetupRouter()
	r.Run(":2019")
}
