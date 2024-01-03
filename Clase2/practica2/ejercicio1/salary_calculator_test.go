package salary_calculator_test

import (
	salary_calculator "gobases/Clase2/practica2/ejercicio1"
	"testing"
)

func TestSalary_calculator(t *testing.T) {
	t.Run("Menos de 50000", func(t *testing.T) {
		var salary int = 40000
		var result int = salary_calculator.Salary_calculator(salary)

		var expected int = 0
		if expected != result {
			t.Errorf("Hubo un error, resultado %d y esperado %d", result, expected)
		}
	})

	t.Run("Mas de 50000", func(t *testing.T) {
		var salary int = 60000
		var result int = salary_calculator.Salary_calculator(salary)

		var expected int = 17
		if expected != result {
			t.Errorf("Hubo un error, resultado %d y esperado %d", result, expected)
		}
	})

	t.Run("Mas de 120000", func(t *testing.T) {
		var salary int = 200000
		var result int = salary_calculator.Salary_calculator(salary)

		var expected int = 27
		if expected != result {
			t.Errorf("Hubo un error, resultado %d y esperado %d", result, expected)
		}

	})

}
