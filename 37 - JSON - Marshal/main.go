package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cao := cachorro{
		Nome:  "Belinha",
		Raca:  "Vira Lata",
		Idade: 10,
	}

	// a função marshal converte uma struct ou map em um slice de bytes para converter en json
	caoJSON, erro := json.Marshal(cao)

	if erro != nil {
		log.Fatal(erro)
	}

	// a função newBuffer converte o slice de bytes devolvido pela função marshal em caracteres
	fmt.Println(bytes.NewBuffer(caoJSON))

	usuario := map[string]string{
		"nome":      "Wiris",
		"sobrenome": "Wernek",
	}

	usuarioJSON, erro := json.Marshal(usuario)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(usuarioJSON))

}
