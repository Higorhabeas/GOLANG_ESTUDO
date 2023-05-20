package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//criando a variável do tipo WaitGroup do pacote sync
	var waitGroup sync.WaitGroup
	//informando a quantidade de go routines
	waitGroup.Add(2)
	//funão anônima para escrevar com concorrência e avisando que Done ti
	go func() {
		escrever("Olá mundo!")
		waitGroup.Done() //-1
	}()

	go func() {
		escrever("Programando em GO!")
		waitGroup.Done() //-1
	}()

	waitGroup.Wait() //esperando o metodo Done() ser executado e Add passar de 2 para zero

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
