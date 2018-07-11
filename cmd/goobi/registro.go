package main

import (
	"github.com/ibraimgm/goobi/internal/app/obi2017"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

func registrarTudo() map[string]*obiinfo.OBIInfo {
	exercicios := make(map[string]*obiinfo.OBIInfo)

	registrar(exercicios, obi2017.Cofre)
	registrar(exercicios, obi2017.Teleferico)

	return exercicios
}

func registrar(exercicios map[string]*obiinfo.OBIInfo, info *obiinfo.OBIInfo) {
	exercicios[info.Nome] = info
}
