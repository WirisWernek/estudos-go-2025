package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// o (u usuario) serve para acoplar a função a estrutura e permitir que a função seja invocada por qualquer variável do mesmo tipo
// o (u usuario) também permite que os valores da variável sejam acessível pela função sem a necessidade de parametros a mais
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s com a idade %d \n", u.nome, u.idade)
}

// assim como funções normais também é possível retornar valores
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// para alterar os valores da variável é necessário declarar o tipo como ponteiro, porém dentro da função não é necessária a desreferenciação, ela pode ser usada e acessada como uma variável normal
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	meuUsuario := usuario{"Wiris", 22}
	fmt.Println(meuUsuario)
	fmt.Println(meuUsuario.maiorDeIdade())
	meuUsuario.salvar()

	segundoUsuario := usuario{"Marcos", 17}
	fmt.Println(segundoUsuario)
	fmt.Println(segundoUsuario.maiorDeIdade())
	segundoUsuario.salvar()

	segundoUsuario.fazerAniversario()
	fmt.Println(segundoUsuario.idade)
}
