package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	altura  float64
	largura float64
}

type circle struct {
	raio float64
}

// Aqui temos assinaturas de métodos.
type forma interface {
	// Crio o método área que retorna um float64.
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %.2f \n", f.area())
}

// Para que eu consiga usar o método da interface, preciso criar o método do struct.
func (rec rectangle) area() float64 {
	return rec.altura * rec.largura
}

func (cir circle) area() float64 {
	return math.Pi * math.Pow(cir.raio, 2)
}

func main() {
	retangulo := rectangle{10, 10}
	circulo := circle{10}
	escreverArea(retangulo)
	escreverArea(circulo)
}
