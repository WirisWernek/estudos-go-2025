package formas_test

import (
	"math"
	. "testes-avancados/formas"
	"testing"
)

func TestArea(t *testing.T) {

	// o t.Run serve para subdividir o teste em diferentes aplicações dele, no caso o teste de area é aplicado ao calculo tanto do retângulo quanto do circulo
	t.Run("Retângulo", func(t *testing.T) {
		retangulo := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := retangulo.Area()

		if areaEsperada != areaRecebida {
			// quando um erro é lançado os demais testes continuam sendo executados
			// t.Errorf("A área recebida %.2f é diferente da area esperada %.2f !", areaRecebida, areaEsperada)

			// quando um fatal é lançado a execução dos testes é parada
			t.Fatalf("A área recebida %.2f é diferente da area esperada %.2f !", areaRecebida, areaEsperada)
		}
	})
	t.Run("Círculo", func(t *testing.T) {

		circulo := Circulo{10}
		areaEsperada := float64(math.Pi * math.Pow(circulo.Raio, 2))
		areaRecebida := circulo.Area()

		if areaEsperada != areaRecebida {
			// quando um erro é lançado os demais testes continuam sendo executados
			// t.Errorf("A área recebida %.2f é diferente da area esperada %.2f !", areaRecebida, areaEsperada)

			// quando um fatal é lançado a execução dos testes é parada
			t.Fatalf("A área recebida %.2f é diferente da area esperada %.2f !", areaRecebida, areaEsperada)
		}

	})

}
