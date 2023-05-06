package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel2, variavel1)

	//PONTEIRO É UMA REFERÊNCA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int //o * indica o endereço de memória

	variavel3 = 222
	ponteiro = &variavel3 // o & indca que é para buscar o endereço da memória

	fmt.Println(variavel3, ponteiro)
	fmt.Println(*ponteiro)

	variavel3 = 333
	fmt.Println(*ponteiro)

}
