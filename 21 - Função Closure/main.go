package main

import "fmt"

// a ideia do closure é basicamente o escopo da função, mesmo que você chame um função que possui variáveis com o mesmo nome da que realizou a chamada, o go segrega essas variáveis utilizando a do escopo da função e não de forma global
func closure() func() {
	texto := "Dentro do closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da main"

	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
