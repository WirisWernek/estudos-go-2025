package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	var posicao uint = 10

	// valor em um determinada posição
	fmt.Println(fibonacci(posicao))

	// sequência ate a posição
	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
