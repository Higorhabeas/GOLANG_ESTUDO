package main

import (
	"fmt"
	"time"
)

func main() {
	//criando o canal
	canal := make(chan string)
	//chamando a função escrever com concorrência
	go escrever("Olá mundo!", canal)

	fmt.Println("Depois da função escrever!")

	/*for {
		//além das informações o canal recebe um booleano informando se aberto ou fechado
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}*/
	/*esta é uma forma mais simples de trafegar pelo canal enquanto ele
	estiver aberto*/
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do loop e do programa!!!")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		//fmt.Println(texto)
		time.Sleep(time.Second)
	}

	close(canal)
}
