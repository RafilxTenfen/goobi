package main

import (
	"github.com/ibraimgm/goobi/internal/app/obi2017"
	"github.com/ibraimgm/goobi/internal/app/obi2018"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

func registrarTudo() map[string]*obiinfo.OBIInfo {
	exercicios := make(map[string]*obiinfo.OBIInfo)

	registrar(exercicios, obi2018.Escadinha)
	registrar(exercicios, obi2018.Piramide)
	registrar(exercicios, obi2018.Xadrez)
	registrar(exercicios, obi2017.Cofre)
	registrar(exercicios, obi2017.Teleferico)
	registrar(exercicios, obi2017.Castelos)
	registrar(exercicios, obi2017.Tabuleiro)
	registrar(exercicios, obi2017.Montanha)

	return exercicios
}

func registrar(exercicios map[string]*obiinfo.OBIInfo, info *obiinfo.OBIInfo) {
	exercicios[info.Nome] = info
}
