package main

import "fmt"

var year map[string]int = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

func month(month string) {
	fmt.Println("El mes ingresado es el mes numero", year[month])
}
func main() {
	month("July")
}
