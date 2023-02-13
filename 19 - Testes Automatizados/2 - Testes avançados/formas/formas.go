package formas

import (
	"math"
)

type FormaGeometrica interface {
	Area() float64
}

type Retangulo struct {
	Altura , Largura float64
}

type Circulo struct {
	Raio float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}
