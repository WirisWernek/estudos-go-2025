package main

import "fmt"

// a generic é basicamente uma interface sem nenhuma definição
// só deve ser usada casos onde uma função realmente precise receber qualquer tipo e não como forma de burlar a tipagem
func generica(generic interface{}) {
	fmt.Println(generic)
}

func main() {
	generica("")
	generica(1)
	generica(true)

	// exemplo de um uso onde o tipo genérico se torna uma gambiarra
	mapa := map[interface{}]interface{}{
		1:            "Texto",
		float32(100): true,
		true:         1,
		"string":     "string",
	}

	fmt.Println(mapa)
}
