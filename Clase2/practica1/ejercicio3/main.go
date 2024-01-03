package main

import "fmt"

func main() {
	fmt.Printf("El salario que recibe es de %f\n", salaryCalculator(160000, "A"))
}

func salaryCalculator(minutes int, category string) (salary float32) {
	var hours int = minutes / 60

	switch category {
	case "A":
		salary = float32(hours * 3000)
		return salary * 1.50
	case "B":
		salary = float32(hours * 1500)
		return salary * 1.20
	case "C":
		return float32(hours) * 1000

	}
	return 27
}
