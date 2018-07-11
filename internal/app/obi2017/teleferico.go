package obi2017

import (
	"github.com/ibraimgm/goobi/internal/iohandler"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

// Teleferico é a solução para o exercício 'teleferico' da OBI 2017/1
var Teleferico = obiinfo.New("teleferico", 2017, 1, "https://olimpiada.ic.unicamp.br/pratique/p1/2017/f1/teleferico/", teleferico)

func teleferico(io iohandler.IOHandler) {
	capacidade := io.ReadInt()
	alunos := io.ReadInt()

	temp := alunos / (capacidade - 1)
	viagens := int(temp)

	if (viagens * (capacidade - 1)) < alunos {
		viagens++
	}

	io.WriteInt(viagens)
}
