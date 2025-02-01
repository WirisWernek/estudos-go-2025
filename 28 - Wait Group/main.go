package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// a estratégia de uso do waitgroups consistem em executar múltiplas goroutines de forma que o programa só termine quando todas elas terminarem. é como se vc tivesse várias funções async em js e o programa só imprime o resultado quando um laço for é executado e garante que todas essas funções já obtiveram seu retorno
	var waitGroup sync.WaitGroup

	waitGroup.Add(3)

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done()
	}()

	go func() {
		escrever("Brincando com Goroutines")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
