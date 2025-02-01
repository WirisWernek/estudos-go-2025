package main

import "fmt"

var n int

// pode ser utilizada em cada arquivo sem a restrição de uma por pacote
// tipo um ngOnInit do Angular
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}
