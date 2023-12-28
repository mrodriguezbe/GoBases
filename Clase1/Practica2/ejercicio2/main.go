package main

import "fmt"

func verifyClient(age int, currentWorking bool, seniority int, salary int) {
	if age > 22 && currentWorking && seniority > 1 {
		if salary > 100000 {
			fmt.Println("Prestamo SIN intereses")
		} else {
			fmt.Println("Prestamo CON intereses")
		}
	} else {
		fmt.Println("NO hay prestamo")
	}
}

func main() {
	verifyClient(23, false, 3, 90000)
}
