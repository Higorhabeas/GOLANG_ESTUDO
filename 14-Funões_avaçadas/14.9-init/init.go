package main

import "fmt"

var n int

func init() {
	n = 25
	fmt.Println("Executando a função init")
}

func main() {
	fmt.Println("Função main")

	fmt.Println(n)
}
