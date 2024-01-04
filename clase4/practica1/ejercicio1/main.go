package main

import (
	"fmt"
)

func (e *salaryError) Error() string {
	return e.msgError
}

type salaryError struct {
	msgError string
}

func main() {
	var salary int = 500

	var err error = &salaryError{
		msgError: "Error: the salary entered does not reach the taxable minimum",
	}

	if salary < 150000 {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("You must pay the tax")
}
