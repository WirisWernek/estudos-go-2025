package main

import "fmt"

func main() {
	// Função anônima
	func() {
		fmt.Println("Função anônima")
	}()

	// Função anônima com parâmetro e retorno
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Função anônima com parâmetro")

	fmt.Println(retorno)
}
