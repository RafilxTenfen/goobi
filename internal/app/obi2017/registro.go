package obi2017

import "github.com/ibraimgm/goobi/internal/obiinfo"

// Registrar registra as soluções disponíveis
func Registrar(lista []obiinfo.OBIInfo) []obiinfo.OBIInfo {
	return append(lista, obiinfo.New("teleferico", 2017, 1, teleferico))
}
