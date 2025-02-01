package main

import "fmt"

func main() {
	// utilizando o buffer você pode dizer o limite de mensagens que o seu canal suporta receber sem a necessidade de travar o processamento do programa, ou seja até ele atingir o limite ele pode executar as demais ações sem causar um deadlock, assim permitindo ser usado em uma mesma função
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Brincando com Buffer"

	close(canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}
