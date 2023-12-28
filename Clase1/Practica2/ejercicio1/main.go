package main

import "fmt"

func letterCounter(word string) {
	fmt.Println(len(word))
	for _, letter := range word {
		fmt.Print(string(letter), " ")
	}
	fmt.Print("\n")
}

func main() {
	letterCounter("estequiometrico")
}
