package main

import (
	"errors"
	"fmt"
)

func main() {
	//existem int8, int32, int64
	var numero int16 = -100
	fmt.Println(numero)

	//uint são para números positivos
	var numero2 uint32 = 200
	fmt.Println(numero2)

	//alias para int32 rune
	var numero3 rune = 1246
	fmt.Println(numero3)

	//alias para uint8 é byte
	var numero4 byte = 35
	fmt.Println(numero4)

	//existem o float32 e float64
	var numeroreal1 float32 = 123.35
	fmt.Println(numeroreal1)

	var numeroreal2 float64 = 325.13
	fmt.Println(numeroreal2)

	/*pela inferencia de tipo o golang coloca o número como float, mas para declarar é necessário expecificar
	se é 32 ou 64*/
	numeroreal3 := 16655.44
	fmt.Println(numeroreal3)

	//strings da mesma forma e uma curosidade é que no go não tem char
	var str string = "texto"
	fmt.Println(str)

	//se vc colocar a aspas simples ' ' o go vai gerar o nº ASCII
	char := 'B'
	fmt.Println(char)

	// o padrão para boleano é false
	var booleano1 bool
	fmt.Println(booleano1)

	//para o go erro é uma tipo e o valor zero dele é nil
	var erro error
	fmt.Println(erro)

	//gerando um erro com o pacote nativo errors
	erro2 := errors.New("Erro interno")
	fmt.Println(erro2)
}
