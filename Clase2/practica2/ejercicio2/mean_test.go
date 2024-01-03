package mean_test

import (
	mean "gobases/Clase2/practica2/ejercicio2"
	"testing"
)

func TestMean(t *testing.T) {
	t.Run("Mas de 120000", func(t *testing.T) {
		var result int = mean.Mean(4, 5, 6, 7, 8)

		var expected int = 6
		if expected != result {
			t.Errorf("Hubo un error, resultado %d y esperado %d", result, expected)
		}

	})

}
