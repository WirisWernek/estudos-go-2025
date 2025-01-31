package main

import "fmt"

// Função com parâmetros variáveis
func soma(numeros ...int) (total int) {
	total = 0
	for _, numero := range numeros {
		total += numero
	}
	return
}

// Função com parâmetros variáveis e parâmetros fixos
// O parâmetro variável deve ser o último
// O parâmetro variável é entendido pela função como um slice
// é possível utilizar um parâmetro fixo e um variável na mesma função desde que o variável seja o último
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalSoma)
	escrever("Olá", 1, 2, 3, 4, 5)
}
