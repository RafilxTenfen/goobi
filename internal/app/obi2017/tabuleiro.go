package obi2017

import (
	"github.com/ibraimgm/goobi/internal/iohandler"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

const (
	celulaBranca = 0
	celulaPreta  = 1
	celulaVazia  = 9
)

// Tabuleiro é a solução para o exercício 'tabuleiro' da OBI 2017/2
var Tabuleiro = obiinfo.New("tabuleiro", 2017, 2, "https://olimpiada.ic.unicamp.br/pratique/p1/2017/f2/tabuleiro/", tabuleiro)

func tabuleiro(io iohandler.IOHandler) {
	n := io.ReadInt()
	board := make([]int, n*n)

	for i := range board {
		board[i] = io.ReadInt()
	}

	io.WriteInt(computeCell(board, n, n, n))
}

func computeCell(board []int, n, i, j int) int {
	if v := getCell(board, n, i, j); v != celulaVazia {
		return v
	}

	a := computeCell(board, n, i, j-1)
	b := computeCell(board, n, i-1, j-1)
	c := computeCell(board, n, i-1, j)

	var cor int
	if (a + b + c) >= celulaPreta {
		cor = celulaBranca
	} else {
		cor = celulaPreta
	}

	setCell(board, n, i, j, cor)
	return cor
}

func getCell(board []int, n, i, j int) int {
	return board[n*(i-1)+j-1]
}

func setCell(board []int, n, i, j, cor int) {
	board[n*(i-1)+j-1] = cor
}
