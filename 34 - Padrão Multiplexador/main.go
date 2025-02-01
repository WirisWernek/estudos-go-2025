package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("By 1"), escrever("By 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// A ideia do multiplexar é unir dois canais em um único
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	saida := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				saida <- mensagem

			case mensagem := <-canalDeEntrada2:
				saida <- mensagem
			}
		}
	}()

	return saida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
