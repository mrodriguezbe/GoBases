package main

import "fmt"

func main() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30}
	fmt.Println("La edad de Benjamin es", employees["Benjamin"])

	var adults int
	for _, age := range employees {
		if age > 21 {
			adults++
		}
	}
	fmt.Println("Los empleados mayores a 21 son", adults)

	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)
}
