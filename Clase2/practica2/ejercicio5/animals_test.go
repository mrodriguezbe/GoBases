package animals_test

import (
	animals "gobases/Clase2/practica2/ejercicio5"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFoodCalculator(t *testing.T) {
	t.Run("Testing with dogs", func(t *testing.T) {
		var animal string = "dog"
		var amount float32 = 10
		result, err := animals.FoodCalculator(animal, amount)
		var expected float32 = 100
		var expected_err = ""

		require.Equal(t, expected_err, err)
		require.Equal(t, expected, result)
	})

	t.Run("Testing with cats", func(t *testing.T) {
		var animal string = "cat"
		var amount float32 = 10
		result, err := animals.FoodCalculator(animal, amount)
		var expected float32 = 50.0
		var expected_err = ""

		require.Equal(t, expected_err, err)
		require.Equal(t, expected, result)
	})

	t.Run("Testing with hamsters", func(t *testing.T) {
		var animal string = "hamster"
		var amount float32 = 10
		result, err := animals.FoodCalculator(animal, amount)
		var expected float32 = 2.5
		var expected_err = ""

		require.Equal(t, expected_err, err)
		require.Equal(t, expected, result)
	})

	t.Run("Testing with spiders", func(t *testing.T) {
		var animal string = "spider"
		var amount float32 = 10
		result, err := animals.FoodCalculator(animal, amount)
		var expected float32 = 1.25
		var expected_err = ""

		require.Equal(t, expected_err, err)
		require.Equal(t, expected, result)
	})

}
