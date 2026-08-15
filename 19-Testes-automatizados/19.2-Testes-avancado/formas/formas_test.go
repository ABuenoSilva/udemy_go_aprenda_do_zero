package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 20}
		areaEsperada := float64(200)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Retângulo - A área recebida %f é diferente da esperada %f.", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Círculo - A área recebida %f é diferente da esperada %f.", areaRecebida, areaEsperada)
		}
	})
}
