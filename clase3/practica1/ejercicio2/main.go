package main

import "fmt"

type Person struct {
	ID          string
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       string
	Position string
	Person   Person
}

func (e Employee) Print() {
	fmt.Println(e)
}

func main() {
	Chris := Person{
		ID:          "32445678",
		Name:        "Chris",
		DateOfBirth: "4/5/1990",
	}

	Zerozeroone := Employee{
		ID:       "001",
		Position: "Developer",
		Person:   Chris,
	}

	Zerozeroone.Print()
}
