package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")
	auxiliar.Escrever()
	fmt.Println(checkmail.ValidateFormat("wiriswernek@"))
	fmt.Println(checkmail.ValidateFormat("wiriswernek@gmail.com"))
}
