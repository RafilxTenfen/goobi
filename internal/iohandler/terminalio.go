package iohandler

import "fmt"

type terminalIO struct{}

func (io terminalIO) ReadInt() int {
	var value int
	fmt.Scan(&value)
	return value
}

func (io terminalIO) WriteInt(value int) {
	fmt.Printf("%d", value)
}

//GetTerminalIO retorna um novo IOHandler para o terminal atual
func GetTerminalIO() IOHandler {
	return terminalIO{}
}
