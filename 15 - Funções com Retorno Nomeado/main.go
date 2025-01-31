package main

import "fmt"

// ao nomear a variável de retorno, não é necessário utilizar o return especificado as variáveis de retorno nem declarar as variáveis de retorno
func calculosMatematicos(n1 int, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
