package obi2017

import (
	"github.com/ibraimgm/goobi/internal/iohandler"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

// Cofre é a solução para o exercício 'cofre' da OBI 2017/1
var Cofre = obiinfo.New("cofre", 2017, 1, "https://olimpiada.ic.unicamp.br/pratique/p1/2017/f1/cofre/", cofre)

func cofre(io iohandler.IOHandler) {
	tamanhoBarra := io.ReadInt()
	tamanhoPosicoes := io.ReadInt()

	barra := make([]int, tamanhoBarra)
	for i := 0; i < tamanhoBarra; i++ {
		barra[i] = io.ReadInt()
	}

	totais := [10]int{}

	from := io.ReadInt()
	totais[barra[from-1]]++

	for i := 1; i < tamanhoPosicoes; i++ {
		to := io.ReadInt()
		var slice []int

		if from < to {
			slice = barra[from:to]
		} else {
			slice = barra[to-1 : from-1]
		}

		for _, number := range slice {
			totais[number]++
		}

		from = to
	}

	for i := 0; i < 10; i++ {
		io.WriteInt(totais[i])
	}
}
