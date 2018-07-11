package obiinfo

import (
	"fmt"

	"github.com/ibraimgm/goobi/internal/iohandler"
)

// OBIFunc é a assinatura usada para as funções referentes aos exercícios da OBI
type OBIFunc func(handler iohandler.IOHandler)

// OBIInfo é uma estrutura com as informações de determinado exercício da OBI
type OBIInfo struct {
	Nome string
	Ano  uint
	Fase uint
	URL  string
	Run  OBIFunc
}

// New retorna uma nova instância de OBIInfo
func New(nome string, ano, fase uint, url string, run OBIFunc) *OBIInfo {
	return &OBIInfo{nome, ano, fase, url, run}
}

// PrintHeader imprime um cabecalho com as informacoes do exercício
func (info *OBIInfo) PrintHeader() {
	fmt.Printf("--\nExercício: %s (%d, fase %d)\nURL: %s\n--\n", info.Nome, info.Ano, info.Fase, info.URL)
}
