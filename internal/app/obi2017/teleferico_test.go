package obi2017

import (
	"testing"

	"github.com/ibraimgm/goobi/internal/iohandler"
)

func TestTeleferico_Ex1(t *testing.T) {
	telefericoTestHelper(t, 10, 20, 3)
}

func TestTeleferico_Ex2(t *testing.T) {
	telefericoTestHelper(t, 12, 55, 5)
}

func TestTeleferico_Ex3(t *testing.T) {
	telefericoTestHelper(t, 100, 87, 1)
}

func telefericoTestHelper(t *testing.T, c, a, esperado int) {
	handler := iohandler.GetBufferedIO()

	handler.InputInt(c)
	handler.InputInt(a)
	teleferico(handler)

	result := handler.OutputInt()

	if result != esperado {
		t.Errorf("Resultado: %d, esperado %d.", result, esperado)
	}
}
