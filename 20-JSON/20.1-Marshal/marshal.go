package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // Se quiser que um campo não seja convertido pode utilizar "-" na tag
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal("Erro na conversão do json: ", erro)
	}

	fmt.Println(cachorroEmJson)
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	c2 := map[string]string{"nome": "Toby", "raca": "poodle"}
	cachorro2EmJson, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal("Erro na conversão do json: ", erro)
	}
	fmt.Println(cachorro2EmJson)
	fmt.Println(bytes.NewBuffer(cachorro2EmJson))
}
