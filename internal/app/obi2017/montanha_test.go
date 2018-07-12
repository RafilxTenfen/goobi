package obi2017

import (
	"testing"

	"github.com/ibraimgm/goobi/internal/iohandler"
)

func TestMontanha_Ex1(t *testing.T) {
	montanhaTestHelper(t, 8, []int{2, 3, 5, 6, 7, 5, 4, 2}, "N")
}

func TestMontanha_Ex2(t *testing.T) {
	montanhaTestHelper(t, 8, []int{2, 3, 6, 5, 4, 6, 3, 2}, "S")
}

func montanhaTestHelper(t *testing.T, n int, picos []int, esperado string) {
	handler := iohandler.GetBufferedIO()

	handler.InputInt(n)
	for _, v := range picos {
		handler.InputInt(v)
	}

	montanha(handler)

	x := handler.OutputStr()

	if x != esperado {
		t.Errorf("Resultado: %q, esperado %q.", x, esperado)
	}
}
