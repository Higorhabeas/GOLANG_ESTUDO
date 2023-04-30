package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "veriavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "varialvel 5", "variaval 6"

	fmt.Println(variavel5, variavel6)

	const constant1 string = "Constante 1"

	fmt.Println(constant1)
}
