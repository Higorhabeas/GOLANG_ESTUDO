package main

import "fmt"

func diadasemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}

}

func diadasemana2(Numero int) string {
	var diasemana string
	switch {
	case Numero == 1:
		diasemana = "Domingo"
	case Numero == 2:
		diasemana = "Segunda-Feira"
	case Numero == 3:
		diasemana = "Terça-Feira"
	case Numero == 4:
		diasemana = "Quarta-Feira"
	case Numero == 5:
		diasemana = "Quinta-Feira"
	case Numero == 6:
		diasemana = "Sexta-Feira"
	case Numero == 7:
		diasemana = "Sábado"
	default:
		diasemana = "Número Inválido"
	}
	return diasemana

}

func main() {
	fmt.Println("Switch")

	dia := diadasemana(200)
	dia2 := diadasemana2(5)
	fmt.Println(dia)
	fmt.Println(dia2)

}
