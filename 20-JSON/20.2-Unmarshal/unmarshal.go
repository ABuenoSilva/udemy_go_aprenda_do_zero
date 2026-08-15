package main

import (
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
	cachorroEmJson := `{"nome":"Rex","raca":"Dálmata","idade":3}`
	var c cachorro
	fmt.Println(c)
	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal("Falha no Unmarshal: ", erro)
	}
	fmt.Println(c)

	cachorro2EmJson := `{"nome":"Toby","raca":"Poodle"}`
	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal("Falha no Unmarshal: ", erro)
	}
	fmt.Println(c2)
}
