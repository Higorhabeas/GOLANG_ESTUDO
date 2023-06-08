package main

import (
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
	cachorroEmjson := `{"nome":"Rex","reca":"Dálmata","idade":3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmjson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2Emjson := `{"nome":"Sorin", "raca":"Labrador"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2Emjson), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)

}
