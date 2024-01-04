// para leerlo

package main

import (
	"fmt"
	"os"
)

func ReadFile(path string) (*os.File, error) {
	return os.Open(path)
}

func main() {
	defer fmt.Println("Ejecución terminada")

	_, err := ReadFile("archivo.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged.")
	}
}
