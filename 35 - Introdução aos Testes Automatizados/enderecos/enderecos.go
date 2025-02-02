package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Verifica se o endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "rodovia", "avenida", "estrada"}

	primeiraPalavara := strings.Split(endereco, " ")[0]
	enderecoValido := false

	for _, tipo := range tiposValidos {
		if strings.EqualFold(tipo, strings.ToLower(primeiraPalavara)) {
			enderecoValido = true
		}
	}

	if enderecoValido {
		caser := cases.Title(language.BrazilianPortuguese)
		return caser.String(primeiraPalavara)
	}

	return "Tipo Inválido"

}
