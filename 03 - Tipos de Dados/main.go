package main

import (
	"errors"
	"fmt"
)

func main() {
	// inteiros
	// int8, int16, int32(rune), int64, int(padrão do sistema, dessa forma o int assume o valor da arquitetura do computador. Ex: arm64 = int64, x32 = int32)

	var numero1 int8 = 127
	var numero2 int = -128
	fmt.Println(numero1)
	fmt.Println(numero2)

	// inteiros positivos(unsygned)
	// uint8(byte), uint16, uint32, uint64, uint(padrão do sistema)

	var numero3 uint32 = 55
	var numero4 rune = 4566
	var numero5 byte = 124

	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero5)

	// floats
	// float32, float64
	// nao pode ser declarado somente "float" pode ser declarado por inferência :=, dessa forma o float assume o valor da arquitetura do computador: Ex numero := 121.21111 , assume que é um tipo float

	var numero6 float32 = 124.44545456465
	fmt.Println(numero6)

	var numero7 float64 = 87874488.1464454545876984
	fmt.Println(numero7)

	numero8 := 48546946489.4565465456465446
	fmt.Println(numero8)

	// strings
	// string

	var str1 string = "variavel tipo string"
	fmt.Println(str1)

	str2 := "variavel string ehhh"
	fmt.Println(str2)

	// caracter
	// char
	// por padrão o char nao existe no go, a declaração em aspas simples retorna o numero do carácter na tabela ASCC, só aceita um carácter e é convertido a um int
	caracter := 'b'
	fmt.Println(caracter)

	// Conceito valor zero
	// variáveis declaradas mas n inicializadas com um valor assumem o valor zero do tipo da variável. Ex: string = " ", int = 0, boll = false, float = 0, error = nil

	var str3 string
	fmt.Println(str3)

	var numero9 int
	fmt.Println(numero9)

	// bool
	var booleano1 bool = false
	fmt.Println(booleano1)

	var booleano2 bool = true
	fmt.Println(booleano2)

	// error
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
