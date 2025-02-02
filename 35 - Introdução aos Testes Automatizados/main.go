package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoDeEndereco)

	tipoDeEndereco = enderecos.TipoDeEndereco("Praça da Sé")
	fmt.Println(tipoDeEndereco)

	tipoDeEndereco = enderecos.TipoDeEndereco("Estrada 284")
	fmt.Println(tipoDeEndereco)
}
