package main

import "fmt"

func closure() func() {

	texto := "Dentro da funçao texto"

	funcao := func() {
		fmt.Println(texto)

	}
	return funcao
}

func main() {
	texto := "Dentro da funçao main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
