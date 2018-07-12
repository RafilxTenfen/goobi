package iohandler

import "fmt"

type terminalIO struct {
	newLine bool
}

func (io *terminalIO) ReadInt() int {
	var value int
	fmt.Scan(&value)
	return value
}

func (io *terminalIO) WriteInt(value int) {
	io.printSpaceIfNeeded()
	fmt.Printf("%d", value)
	io.newLine = false
}

func (io *terminalIO) WriteStr(value string) {
	io.printSpaceIfNeeded()
	fmt.Printf("%s", value)
	io.newLine = false
}

func (io *terminalIO) printSpaceIfNeeded() {
	if !io.newLine {
		fmt.Print(" ")
	}
}

//GetTerminalIO retorna um novo IOHandler para o terminal atual
func GetTerminalIO() IOHandler {
	return &terminalIO{true}
}
