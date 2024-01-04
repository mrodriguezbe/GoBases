package main

import (
	"fmt"
)

func main() {
	var salary int = 500

	if salary < 10000 {
		err := fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
		fmt.Println(err.Error())
		return
	}

	fmt.Println("You must pay the tax")
}
