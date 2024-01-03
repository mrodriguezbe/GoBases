package animals

import "fmt"

type CounterFunction = func(float32) float32

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func FoodCalculator(animal string, amount float32) (float32, string) {
	animalFunc, err := Animal(animal)
	return animalFunc(amount), err
}

func Animal(typeAnimal string) (function CounterFunction, err string) {
	switch typeAnimal {
	case dog:
		return func(amount float32) float32 { return amount * 10.0 }, err
	case cat:
		return func(amount float32) float32 { return amount * 5.0 }, err
	case spider:
		return func(amount float32) float32 { return amount * 0.125 }, err
	case hamster:
		return func(amount float32) float32 { return amount * 0.25 }, err
	default:
		return function, "The animal doest exist"
	}
}

func main() {
	animalDog, msg := Animal(dog)
	fmt.Println(msg)
	animalCat, msg := Animal(cat)
	fmt.Println(msg)

	var amount float32

	amount += animalDog(10)
	fmt.Println(amount)

	amount += animalCat(10)
	fmt.Println(amount)
}
