package obi2017

import "github.com/ibraimgm/goobi/internal/iohandler"

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
