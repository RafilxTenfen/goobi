package obiinfo

import "github.com/ibraimgm/goobi/internal/iohandler"

// OBIFunc é a assinatura usada para as funções referentes aos exercícios da OBI
type OBIFunc func(handler iohandler.IOHandler)

// OBIInfo é uma estrutura com as informações de determinado exercício da OBI
type OBIInfo struct {
	nome string
	ano  uint
	fase uint
	run  OBIFunc
}

// New retorna uma nova instância de OBIInfo
func New(nome string, ano, fase uint, run OBIFunc) OBIInfo {
	return OBIInfo{nome, ano, fase, run}
}

// Run executa uma solucao
func (info OBIInfo) Run(handler iohandler.IOHandler) {
	info.run(handler)
}
