package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func min(califications ...int) int {
	var min int = 99999999
	for i := 0; i < len(califications); i++ {
		if califications[i] < min {
			min = califications[i]
		}
	}
	return min
}

func max(califications ...int) int {
	var max int = 0
	for i := 0; i < len(califications); i++ {
		if califications[i] > max {
			max = califications[i]
		}
	}
	return max
}

func mean(califications ...int) int {
	var aux int = 0

	for i := 0; i < len(califications); i++ {
		aux += califications[i]
	}
	return aux / len(califications)
}

func operation(function string) (funct func(...int) int, err string) {
	switch function {
	case minimum:
		return min, ""
	case average:
		return mean, ""
	case maximum:
		return max, ""
	}
	return
}

func main() {
	minFunc, err := operation(minimum)
	fmt.Printf(err)
	averageFunc, err := operation(average)
	fmt.Printf(err)
	maxFunc, err := operation(maximum)
	fmt.Printf(err)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Println(minValue, averageValue, maxValue)
}
