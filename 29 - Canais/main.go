package main

import (
	"fmt"
	"time"
)

// canais tem um funcionamento bem similar a uma queue
func main() {
	canal := make(chan string)
	go escrever("Olá mundo", canal)

	fmt.Println("Depois de chamar a função escrever")

	// forma mais simplificada de consumir os canais
	// for mensagem := range canal {
	// 	println(mensagem)
	// }

	// forma mais extensa de consumir os canais
	for {
		// esperando que o canal receber um valor(ele recebe esse valor e devolve para a variável junto com o dado se ele esta ou não aberto)
		mensagem, aberto := <-canal

		// o canal precisa ser fechado mas também é necessário um if para finalizar o loop quando isso ocorrer
		if !aberto {
			fmt.Println("Canal Fechado")
			break
		}

		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)

	for i := 0; i < 5; i++ {
		// mandando um valor para o canal
		canal <- texto
		time.Sleep(time.Second)
	}

	// se o canal não for fechado pode ocorrer um deadlock
	close(canal)
}
