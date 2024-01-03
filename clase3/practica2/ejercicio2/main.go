package main

import (
	"fmt"
)

type Product interface {
	Price() float64
}

type SmallProduct struct {
	price float64
}

func (p SmallProduct) Price() float64 {
	return p.price
}

type MediumProduct struct {
	price       float64
	maintenance float64
}

func (p MediumProduct) Price() float64 {
	return p.price * p.maintenance
}

type LargeProduct struct {
	price, maintenance, shipping float64
}

func (p LargeProduct) Price() float64 {
	return p.price*p.maintenance + p.shipping
}

func ProductFactory(productType string, price float64) Product {
	switch productType {
	case "Small":
		return SmallProduct{price}
	case "Medium":
		return MediumProduct{price, 1.03}
	case "Large":
		return LargeProduct{price, 1.06, 2500.0}
	default:
		panic("Tipo de producto no válido")
	}
}

func main() {
	small := ProductFactory("Small", 100.0)
	medium := ProductFactory("Medium", 100.0)
	large := ProductFactory("Large", 100.0)

	fmt.Println("Precio Small:", small.Price())
	fmt.Println("Precio Medium:", medium.Price())
	fmt.Println("Precio Large:", large.Price())
}
