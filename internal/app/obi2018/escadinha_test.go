package obi2018

import (
	"testing"

	"github.com/ibraimgm/goobi/internal/iohandler"
)

func TestEscadinha_Ex1(t *testing.T) {
	escadinhaTestHelper(t, 8, []int{1, 1, 1, 3, 5, 4, 8, 12}, 4)
}

func TestEscadinha_Ex2(t *testing.T) {
	escadinhaTestHelper(t, 1, []int{112}, 1)
}

func TestEscadinha_Ex3(t *testing.T) {
	escadinhaTestHelper(t, 5, []int{11, -106, -223, -340, -457}, 1)
}

func TestEscadinha_Ex4(t *testing.T) {
	escadinhaTestHelper(t, 2, []int{1, 2}, 1)
}

func escadinhaTestHelper(t *testing.T, n int, sequencia []int, esperado int) {
	handler := iohandler.GetBufferedIO()

	handler.InputInt(n)
	for _, v := range sequencia {
		handler.InputInt(v)
	}

	escadinha(handler)

	x := handler.OutputInt()

	if x != esperado {
		t.Errorf("Resultado: %d, esperado %d.", x, esperado)
	}
}
