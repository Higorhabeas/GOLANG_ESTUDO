package main

import "fmt"

func recuperaraExecução() {
	fmt.Println("Tentativa de recuperar")
	if r := recover(); r != nil {
		fmt.Println("Programa recuperado com sucesso!")
	}
}

func alunoaprovado(n1, n2 float64) bool {
	defer recuperaraExecução()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é igual a 6")
}

func main() {
	fmt.Println("Panic e Recover")

	fmt.Println(alunoaprovado(6, 6))

	fmt.Println("Pós execução!")
}
