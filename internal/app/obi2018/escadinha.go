package obi2018

import (
	"github.com/ibraimgm/goobi/internal/iohandler"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

// Escadinha é a solução para o exercício 'escadinha' da OBI 2018/2
var Escadinha = obiinfo.New("escadinha", 2018, 1, "https://olimpiada.ic.unicamp.br/pratique/p1/2018/f1/escadinha/", escadinha)

func escadinha(io iohandler.IOHandler) {
	escadas := 0
	n := io.ReadInt()
	n1 := io.ReadInt()

	if n > 2 {
		var oldDiff int

		for i := 0; i < n-1; i++ {
			n2 := io.ReadInt()
			diff := abs(n1 - n2)

			if i > 0 {
				if diff != oldDiff {
					escadas++
				}
			}

			oldDiff = diff
			n1 = n2
		}
		escadas++
	} else {
		if n == 2 {
			io.ReadInt() // pela definicao do problema, nao faz diferenca
		}
		escadas = 1
	}

	io.WriteInt(escadas)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
