package calculator_test

import (
	calculator "gobases/Clase2/practica2/ejercicio3"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	t.Run("Empleado tipo A", func(t *testing.T) {
		var minutes int = 1000
		var class = "A"
		var result float32 = calculator.SalaryCalculator(minutes, class)
		var expected float32 = 72000.0

		require.Equal(t, expected, result)
	})

	t.Run("Empleado tipo B", func(t *testing.T) {
		var minutes int = 1000
		var class = "B"
		var result float32 = calculator.SalaryCalculator(minutes, class)
		var expected float32 = 28800.002

		require.Equal(t, expected, result)
	})

	t.Run("Empleado tipo C", func(t *testing.T) {
		var minutes int = 1000
		var class = "C"
		var result float32 = calculator.SalaryCalculator(minutes, class)
		var expected float32 = 16000.0

		require.Equal(t, expected, result)
	})
}
