package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	// quando feito assim mesmo que o canal1 receba uma mensagem a cada 0.5s ele ainda precisa esperar os 1.5s restantes para que o cana2 receba sua mensagem e a interação continue, ou seja isso força o canal1 a ser impresso a cada 2s assim como o canal2
	// for{
	// 	mensagemCanal1 := <- canal1
	// 	fmt.Println(mensagemCanal1)

	// 	mensagemCanal2 := <- canal2
	// 	fmt.Println(mensagemCanal2)
	// }

	for {
		// com o select não é necessário aguardar o recebimento da mensagem, quando ela for recebida ela será impressa
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)

		}
	}
}
