package main

import (
	"fmt"
	"math"
)

// no go não é necessário implementar propriamente a interface
type formaGeometrica interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

// se uma estrutura tem as características definidas em uma interface o go já entende que ela implementa aquela interface
// logo se implementa todos os métodos com a mesma assinatura da interface a estrutura implementou a interface
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	// return math.Pi * (c.raio * c.raio)
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f formaGeometrica) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
