package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Theo",
		"sobrenome": "De bellis",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Renata",
			"ultimo":   "Silva",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"])
	fmt.Println(usuario2["nome"]["primeiro"])
	fmt.Println(usuario2["nome"]["ultimo"])
}
