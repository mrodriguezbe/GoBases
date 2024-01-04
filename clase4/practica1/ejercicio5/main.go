package main

import (
	"errors"
	"fmt"
)

var err error = errors.New("Error: the worker cannot have worked less than 80 hours per month")

func calculator(hours int, price int) (int, error) {
	salary := hours * price

	if hours < 80 || hours < 0 {
		return 0, err
	}

	return salary, nil
}

func main() {
	result, err1 := calculator(90, 100)

	if errors.Is(err1, err) {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}
