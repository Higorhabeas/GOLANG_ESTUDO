package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("recursivas")

	posicao := uint(5)

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
