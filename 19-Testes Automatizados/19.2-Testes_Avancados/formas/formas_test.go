package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Retangulo", func(t *testing.T) {

		ret := Retangulo{10, 12}

		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}

	})

	t.Run("circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)

		areaRecebida := circ.Area()
		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}

	})

}
