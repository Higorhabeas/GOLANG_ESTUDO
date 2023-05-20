package main

import "fmt"

func main() {
	//criando o canal com 2 de buffer
	canal := make(chan string, 2)
	canal <- "Olá mundo!"
	canal <- "Programando em GO!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
