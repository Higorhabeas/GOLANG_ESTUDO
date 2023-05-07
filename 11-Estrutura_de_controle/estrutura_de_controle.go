package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estrutura de controle")

	numero := -10
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Manor que 15")
	}

	if numero2 := numero; numero2 > 0 {
		fmt.Println("Numero é maior que zero")
	} else {
		fmt.Println("Numero menor que zero")
	}

}
