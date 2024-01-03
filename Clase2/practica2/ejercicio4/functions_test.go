package functions_test

import (
	functions "gobases/Clase2/practica2/ejercicio4"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunctions(t *testing.T) {
	t.Run("Max", func(t *testing.T) {
		var result int = functions.Max(4, 5, 6, 7, 8, 9)
		var expected int = 9

		require.Equal(t, expected, result)
	})

	t.Run("Empleado tipo B", func(t *testing.T) {
		var result int = functions.Min(2, 3, 4, 5, 6, 7)
		var expected int = 2

		require.Equal(t, expected, result)
	})

	t.Run("Empleado tipo C", func(t *testing.T) {
		var result int = functions.Mean(2, 3, 4, 5, 6)
		var expected int = 4

		require.Equal(t, expected, result)
	})
}
