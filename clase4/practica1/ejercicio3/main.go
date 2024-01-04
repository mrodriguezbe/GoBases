package main

import (
	"errors"
	"fmt"
)

var err error = errors.New("Error: the salary entered does not reach the taxable minimum")

func x(salary int) error {
	if salary > 10000 {
		return nil
	}
	return err
}

func main() {
	var salary int = 500

	result := x(salary)

	if coincidence := errors.Is(result, err); coincidence {
		fmt.Println(result.Error())
		fmt.Println(coincidence)
		return
	}

	fmt.Println("You must pay the tax")
}
