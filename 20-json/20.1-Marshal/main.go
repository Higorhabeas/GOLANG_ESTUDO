package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"reca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroEmjson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	//fmt.Println(cachorroEmjson)

	fmt.Println(bytes.NewBuffer(cachorroEmjson))

	c2 := map[string]string{
		"nome": "Sorin",
		"raca": "Labrador",
	}

	cachorro2Emjson, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorro2Emjson))

}
