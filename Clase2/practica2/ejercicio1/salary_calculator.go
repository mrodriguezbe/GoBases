package salary_calculator

import "fmt"

func main() {
	fmt.Printf("El impuesto que recibe es de %d\n", Salary_calculator(160000))
}

func Salary_calculator(salary int) (taxes int) {
	if salary < 50000 {
		return 0
	}
	if salary < 150000 {
		return 17
	}
	return 27
}
