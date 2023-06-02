package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func multiplexar(canalentrada, canalentrada2 <-chan string) <-chan string {
	canasaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalentrada:
				canasaida <- mensagem
			case mensagem := <-canalentrada2:
				canasaida <- mensagem
			}
		}
	}()
	return canasaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Int31n(2000)))
		}
	}()

	return canal

}
