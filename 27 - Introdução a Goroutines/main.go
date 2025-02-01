package main

import (
	"fmt"
	"time"
)

func main() {

	// as goroutines servem para dizer que uma função não precisa aguardar seu fim para seguir, pode ser comparada ao conceito de async/await no javascript
	// as goroutines são declaradas pelo prefixo go
	go escrever("Olá Mundo")
	go escrever("Programando em Go")
	escrever("Brincando com Goroutines")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
