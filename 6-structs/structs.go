package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo strucs")

	enderecoexemplo := endereco{"Rua das sempre-vevas", 181}

	var u usuario
	u.nome = "Theo"
	u.idade = 8
	fmt.Println(u)

	//outra forma

	usuario2 := usuario{"Renata", 43, enderecoexemplo}
	fmt.Println(usuario2)

	//outra forma quando não tenho todos os dados
	usuario3 := usuario{nome: "Felipe"}
	fmt.Println(usuario3)
}
