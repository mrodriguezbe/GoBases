package mean

import "fmt"

func main() {
	fmt.Printf("El promedio es de %d\n", Mean(5, 6, 8, 9, 3, 5, 7, 8, 4, 8))
}

func Mean(califications ...int) int {
	var aux int = 0

	for i := 0; i < len(califications); i++ {
		aux += califications[i]
	}
	return aux / len(califications)
}
