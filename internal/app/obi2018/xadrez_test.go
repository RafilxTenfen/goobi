package obi2018

import (
	"testing"

	"github.com/ibraimgm/goobi/internal/iohandler"
)

func TestXadrez_Ex1(t *testing.T) {
	xadrezTestHelper(t, 6, 9, 0)
}

func TestXadrez_Ex2(t *testing.T) {
	xadrezTestHelper(t, 8, 8, 1)
}

func TestXadrez_Ex3(t *testing.T) {
	xadrezTestHelper(t, 5, 91, 1)
}

func TestXadrez_Ex4(t *testing.T) {
	xadrezTestHelper(t, 401, 322, 0)
}

func xadrezTestHelper(t *testing.T, l, c, esperado int) {
	handler := iohandler.GetBufferedIO()
	handler.InputInt(l)
	handler.InputInt(c)

	xadrez(handler)

	x := handler.OutputInt()
	if x != esperado {
		t.Errorf("Resultado: %d, esperado %d.", x, esperado)
	}
}
