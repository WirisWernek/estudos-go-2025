package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"-"` // o - serve para ignorar o valor na conversão
}

func main() {
	cachorroEmJSON := `{"nome":"Belinha","raca":"Vira Lata","idade":10}`

	var cachorro cachorro

	// a função unmarshal mapeia um slice de bytes para uma struct ou map
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &cachorro); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro)

	usuarioEmJSON := `{"nome":"Wiris","sobrenome":"Wernek"}`
	usuario := map[string]string{}

	if erro := json.Unmarshal([]byte(usuarioEmJSON), &usuario); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(usuario)

}
