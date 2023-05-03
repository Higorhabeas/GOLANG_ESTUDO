package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 3 * 2
	divisao := 5 / 2
	resto := 10 % 6
	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

}
