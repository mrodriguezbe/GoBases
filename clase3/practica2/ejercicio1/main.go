package main

import "fmt"

type Student struct {
	Name     string
	Surname  string
	Document string
	Date     string
}

func (s Student) Detail() {
	fmt.Printf("Name: [%s]\n", s.Name)
	fmt.Printf("Surname: [%s]\n", s.Surname)
	fmt.Printf("Document: [%s]\n", s.Document)
	fmt.Printf("Date: [%s]\n", s.Date)
}

func main() {
	maikol := Student{
		Name:     "Maikol",
		Surname:  "Jhonson",
		Document: "9445673",
		Date:     "13/3/2001",
	}

	maikol.Detail()
}
