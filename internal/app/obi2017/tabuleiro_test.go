package obi2017

import (
	"testing"

	"github.com/ibraimgm/goobi/internal/iohandler"
)

func TestTabuleiro_Ex1(t *testing.T) {
	board := []int{
		0, 1,
		1, 9}

	tabuleiroTestHelper(t, 2, board, 0)
}

func TestTabuleiro_Ex2(t *testing.T) {
	board := []int{
		0, 0, 1, 0, 0, 0,
		1, 9, 9, 9, 9, 9,
		0, 9, 9, 9, 9, 9,
		0, 9, 9, 9, 9, 9,
		1, 9, 9, 9, 9, 9,
		1, 9, 9, 9, 9, 9}

	tabuleiroTestHelper(t, 6, board, 1)
}

func tabuleiroTestHelper(t *testing.T, n int, board []int, esperado int) {
	handler := iohandler.GetBufferedIO()
	handler.InputInt(n)

	for _, v := range board {
		handler.InputInt(v)
	}

	tabuleiro(handler)
	x := handler.OutputInt()

	if x != esperado {
		t.Errorf("Resultado: %d, esperado %d.", x, esperado)
	}
}
