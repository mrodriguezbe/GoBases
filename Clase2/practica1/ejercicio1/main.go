package main

import "fmt"

func main() {
	fmt.Printf("El impuesto que recibe es de %d\n", salaryCalculator(160000))
}

func salaryCalculator(salary int) (taxes int) {
	if salary < 50000 {
		return 0
	}
	if salary < 150000 {
		return 17
	}
	return 27
}
