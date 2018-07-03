package main

import "github.com/ibraimgm/goobi/internal/iohandler"

func main() {
	lista := registrarTudo()
	handler := iohandler.GetTerminalIO()

	lista[0].PrintHeader()
	lista[0].Run(handler)
}
