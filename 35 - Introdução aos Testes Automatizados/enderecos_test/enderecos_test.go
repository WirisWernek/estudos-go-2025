package enderecos_test

import (
	. "introducao-testes/enderecos"
	"strings"
	"testing"
)

type cenarioDeTeste struct {
	casoDeTeste     string
	retornoEsperado string
}

// Deve chamar a função com sucesso
func TestTipoDeEndereco(t *testing.T) {
	// serve para indicar que um teste pode ser executado em paralelo com outros com essa notação
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{
			casoDeTeste:     "Rua ABC",
			retornoEsperado: "Rua",
		},
		{
			casoDeTeste:     "Avenida Paulista",
			retornoEsperado: "Avenida",
		},
		{
			casoDeTeste:     "Rodovia dos Imigrantes",
			retornoEsperado: "Rodovia",
		},
		{
			casoDeTeste:     "Praça da Sé",
			retornoEsperado: "Tipo Inválido",
		},
		{
			casoDeTeste:     "Estrada São Gonçalo",
			retornoEsperado: "Estrada",
		},
		{
			casoDeTeste:     "RUA MARECHAL RONDON",
			retornoEsperado: "Rua",
		},
		{
			casoDeTeste:     "",
			retornoEsperado: "Tipo Inválido",
		},
	}

	for _, cenario := range cenariosDeTeste {
		retornoObtido := TipoDeEndereco(cenario.casoDeTeste)

		if !strings.EqualFold(cenario.retornoEsperado, retornoObtido) {
			t.Errorf("O tipo recebido %s é diferente do esperado %s !", cenario.retornoEsperado, retornoObtido)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Teste quebrou")
	}
}
