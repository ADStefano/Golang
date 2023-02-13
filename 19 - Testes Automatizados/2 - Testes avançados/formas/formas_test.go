package formas_test

import (
	."teste-avancado/formas"
	"testing"
	"math"
)

func TestArea(t *testing.T){
	t.Run("Retângulo", func(t *testing.T){
		t.Parallel()
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaRecebida != areaEsperada{
			// Fatalf é igual ao errorf mas ele para o código no erro
			t.Fatalf("Área recebida :%f é diferente da Área esperada: %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T){
		t.Parallel()
		cir := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := cir.Area()

		if areaRecebida != areaEsperada{
			t.Fatalf("Área recebida :%f é diferente da Área esperada: %f", areaRecebida, areaEsperada)
		}
	})
}
