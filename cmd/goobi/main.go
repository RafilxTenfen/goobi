package main

import (
	"fmt"
	"os"

	"github.com/ibraimgm/goobi/internal/iohandler"
	"github.com/pborman/getopt/v2"
)

func main() {
	optHelp := getopt.BoolLong("help", 'h', "Exibe a ajuda")
	optQuiet := getopt.BoolLong("quiet", 'q', "Suprime o cabeçalho do exercício que está sendo executado")
	optList := getopt.BoolLong("list", 'l', "Lista os exercícios disponíveis")
	optRun := getopt.StringLong("run", 'r', "", "Especifica o exercício a ser executado", "teleferico")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		return
	}

	exercicios := registrarTudo()

	if *optList {
		fmt.Printf("Lista dos exercícios disponíveis:\n\n")
		fmt.Printf("%-10s\t%-6s\t%s\n", "Nome", "Ano", "URL")
		for _, e := range exercicios {
			fmt.Printf("%-10s\t%d/%d\t%s\n", e.Nome, e.Ano, e.Fase, e.URL)
		}
		return
	}

	if *optRun == "" {
		fmt.Printf("*** Erro: Não foi especificado nenhum exercício para execução.\nTente '-h' para ver as opções.\n")
		os.Exit(1)
	}

	lista := registrarTudo()
	handler := iohandler.GetTerminalIO()

	if exercicio, exists := lista[*optRun]; exists {
		if !*optQuiet {
			exercicio.PrintHeader()
		}

		exercicio.Run(handler)
	} else {
		fmt.Printf("*** Erro: Não foi encontrado o exercício %q.\nTente '-l' para ver os exercícios disponíveis.\n", *optRun)
		os.Exit(1)
	}
}
