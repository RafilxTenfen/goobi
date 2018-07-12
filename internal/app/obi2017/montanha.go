package obi2017

import (
	"github.com/ibraimgm/goobi/internal/iohandler"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

// Montanha é a solução para o exercício 'montanha' da OBI 2017/2
var Montanha = obiinfo.New("montanha", 2017, 2, "https://olimpiada.ic.unicamp.br/pratique/p1/2017/f2/montanha/", montanha)

func montanha(io iohandler.IOHandler) {
	n := io.ReadInt()

	var a, b, c int
	found := false

	for i := 0; i < n; i++ {
		a = b
		b = c
		c = io.ReadInt()

		if found || a == 0 || b == 0 {
			continue
		}

		if (b < a) && (b < c) {
			found = true
		}
	}

	if found {
		io.WriteStr("S")
	} else {
		io.WriteStr("N")
	}
}
