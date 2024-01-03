package main

import "fmt"

type Product struct {
	ID          string
	Name        string
	Price       float32
	Description string
	Category    string
}

var Products []Product

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	fmt.Println(Products)
}

func GetById(i int) (p Product) {
	if i >= len(Products) {
		fmt.Println("Wrong index bro")
		return p
	}
	fmt.Println(Products[i])
	return Products[i]
}

func main() {
	product1 := Product{
		ID:          "001",
		Name:        "Beer",
		Price:       5.99,
		Description: "Golden",
		Category:    "B",
	}
	product2 := Product{
		ID:          "002",
		Name:        "Coke",
		Price:       3.99,
		Description: "Normal",
		Category:    "A",
	}
	product1.Save()
	product2.Save()
	product1.GetAll()
	GetById(1)
}
