package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {

	defer fmt.Println("Média Calculada Resultado será retornado.")
	fmt.Println("Calculando Média")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	} else {
		return false
	}
}

func main() {
	// serve para adiar a execução de uma função para o ultimo momento possível(antes do return ou no fim da execução)
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(8, 9))
	fmt.Println(alunoEstaAprovado(3, 5))
}
