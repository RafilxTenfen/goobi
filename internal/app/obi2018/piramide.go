package obi2018

import (
	"github.com/ibraimgm/goobi/internal/iohandler"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

// Piramide é a solução para o exercício 'piramide' da OBI 2018/1
var Piramide = obiinfo.New("piramide", 2018, 1, "https://olimpiada.ic.unicamp.br/pratique/p1/2018/f1/piramide/", piramide)

func piramide(io iohandler.IOHandler) {
	n := io.ReadInt()
	caixas := make([]int, n*n)

	for i := range caixas {
		caixas[i] = io.ReadInt()
	}

	io.WriteInt(computeWeight(caixas, n, n, 1, n))
}

func computeWeight(caixas []int, n, line, start, end int) int {
	if start == end {
		return getCell(caixas, n, line, start)
	}

	var sum int
	for i := start; i <= end; i++ {
		sum += getCell(caixas, n, line, i)
	}

	left := computeWeight(caixas, n, line-1, start+1, end)
	right := computeWeight(caixas, n, line-1, start, end-1)

	if left < right {
		return sum + left
	}

	return sum + right
}

func getCell(caixas []int, n, i, j int) int {
	return caixas[n*(i-1)+j-1]
}
