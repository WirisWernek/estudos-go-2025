package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Tentando recuperar Execução")

	// o uso do recover serve para capturar o panic da função e tentar manter a execução posterior ao panic
	if r := recover(); r != nil {
		fmt.Println("Recuperada com sucesso")
	}
}

func calcularMedia(n1, n2 float32) bool {

	defer recuperarExecucao()
	fmt.Println("Calculando Média")

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// o panic serve pra parar a execução, pode ser comparado a um throw new Exception do java sem try catch
	panic("Média é 6")
}

func main() {
	// fmt.Println(calcularMedia(4, 9))
	fmt.Println(calcularMedia(5, 7))
}
