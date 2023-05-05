package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Harença")

	pessoa1 := pessoa{"Theo", "De Bellis", 8, 155}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "3º Ano", "educare"}
	fmt.Println(estudante1)

	fmt.Println(estudante1.curso)

}
