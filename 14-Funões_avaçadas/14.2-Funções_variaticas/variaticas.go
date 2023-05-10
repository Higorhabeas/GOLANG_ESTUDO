package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}

func escrever(text string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(text, numero)
	}
}

func main() {

	totaldasoma := soma(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(totaldasoma)

	escrever("Theo", 1, 4, 2, 4, 5, 7)

}
