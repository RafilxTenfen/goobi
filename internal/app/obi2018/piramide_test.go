package obi2018

import (
	"testing"

	"github.com/ibraimgm/goobi/internal/iohandler"
)

func TestPiramide_Ex1(t *testing.T) {
	caixas := []int{
		5, 2, 4,
		3, 6, 7,
		10, 5, 10,
	}

	piramideTestHelper(t, 3, caixas, 36)
}

func TestPiramide_Ex2(t *testing.T) {
	caixas := []int{
		45, 8, 3, 1,
		1, 10, 5, 67,
		4, 4, 3, 18,
		10, 4, 7, 12,
	}

	piramideTestHelper(t, 4, caixas, 62)
}

func piramideTestHelper(t *testing.T, n int, caixas []int, esperado int) {
	handler := iohandler.GetBufferedIO()
	handler.InputInt(n)

	for _, v := range caixas {
		handler.InputInt(v)
	}

	piramide(handler)

	x := handler.OutputInt()
	if x != esperado {
		t.Errorf("Resultado: %d, esperado %d.", x, esperado)
	}
}
