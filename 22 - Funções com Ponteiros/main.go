package main

import "fmt"

func inverteSinal(numero int) (retorno int) {
	retorno = numero * -1
	numero = retorno
	return
}

func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20

	// utilizando a própria variável é gerada uma cópia da mesma para a função e por isso mesmo que ela seja alterada não reflete na variável passada
	numeroInvertido := inverteSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	// utilizando um ponteiro para a variável qualquer alteração nela será refletido em todas as suas referências
	novoNumero := 10
	inverteSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
