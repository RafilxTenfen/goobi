package obi2017

import (
	"github.com/ibraimgm/goobi/internal/iohandler"
)

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
