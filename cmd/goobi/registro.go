package main

import (
	"github.com/ibraimgm/goobi/internal/app/obi2017"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

func registrarTudo() []obiinfo.OBIInfo {
	lista := make([]obiinfo.OBIInfo, 0)

	lista = obi2017.Registrar(lista)

	return lista
}
