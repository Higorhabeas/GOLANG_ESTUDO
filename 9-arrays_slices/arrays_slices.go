package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//para criar o array colocar o tamanho entre [] e o tipo
	var array1 [5]int
	array1[0] = 200
	fmt.Println(array1)

	array2 := [4]string{"Theo", "Renata", "Felipe", "Higor"}
	fmt.Println(array2)

	//declarar o array com [...] vai fixar de acordo com os dados que vc incluir
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	//no caso acima o array3 está fixado com 5 posições
	array4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array4)
	//no caso acima o array3 está fixado com 10 posições

	slice := []int{-1, -2, 0}

	fmt.Println(slice)
	//verificando o tipo de cada um
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array4))

	// acrescentando itens no slice
	slice = append(slice, 12)
	fmt.Println(slice)

	//criando slice a partir de um array, posição 1 inclusiva e vai até a 4 e ele
	//fonciona como um ponteiro, se ve alterar o array vai refletir no slice
	slice2 := array4[1:5]
	fmt.Println(slice2)

	//arrays internos
	fmt.Println("__________________________")
	//no caso do make se vc estourar o slice a função dobra a capacidade do array
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 20)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 233)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
