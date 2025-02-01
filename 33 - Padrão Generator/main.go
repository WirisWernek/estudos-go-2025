package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá mundo")
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// O objetivo do generator é encapsular todo o processo da goroutine em uma função a parte do processamento principal
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
