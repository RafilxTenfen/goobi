package obi2017

import (
	"testing"

	"github.com/ibraimgm/goobi/internal/iohandler"
)

func TestCofre_Ex1(t *testing.T) {
	cofreTestHelper(t, 14, 5, []int{9, 4, 3, 9, 1, 2, 4, 5, 1, 1, 9, 7, 0, 5}, []int{1, 9, 4, 11, 13}, [10]int{1, 6, 3, 1, 4, 3, 0, 1, 0, 4})
}

func TestCofre_Ex2(t *testing.T) {
	cofreTestHelper(t, 5, 4, []int{5, 8, 0, 5, 1}, []int{1, 4, 2, 5}, [10]int{3, 1, 0, 0, 0, 3, 0, 0, 2, 0})
}

func cofreTestHelper(t *testing.T, n, m int, barra, posicoes []int, esperado [10]int) {
	handler := iohandler.GetBufferedIO()

	handler.InputInt(n)
	handler.InputInt(m)

	for _, v := range barra {
		handler.InputInt(v)
	}

	for _, v := range posicoes {
		handler.InputInt(v)
	}

	cofre(handler)

	for i := 0; i < 10; i++ {
		x := handler.OutputInt()

		if x != esperado[i] {
			t.Errorf("Esperado %d na posição %d (encontrado %d).", esperado[i], i, x)
		}
	}
}
