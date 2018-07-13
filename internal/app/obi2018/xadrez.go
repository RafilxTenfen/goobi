package obi2018

import (
	"github.com/ibraimgm/goobi/internal/iohandler"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

// Xadrez é a solução para o exercício 'xadrez' da OBI 2018/1
var Xadrez = obiinfo.New("xadrez", 2018, 1, "https://olimpiada.ic.unicamp.br/pratique/p1/2018/f1/xadrez/", xadrez)

func xadrez(io iohandler.IOHandler) {
	l := io.ReadInt()
	c := io.ReadInt()

	if l == c || ((l%2 == 0) == (c%2 == 0)) {
		io.WriteInt(1)
	} else {
		io.WriteInt(0)
	}
}
