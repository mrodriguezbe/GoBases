package main

// bytes, err := io.ReadAll(file)
// para transformarlo a string
// string(bytes)

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("ejercicio2/customers.txt")

	if err != nil {
		panic("PANIC OPEN!!!")
	}
	bytes, err := io.ReadAll(file)
	fmt.Println(string(bytes))

	if err != nil {
		panic("Panic READALL!!!")
	}

	defer func() {
		fmt.Println("Ejecución terminada")
		file.Close()
	}()

}
