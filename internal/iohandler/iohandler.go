package iohandler

//IOHandler é uma interface para abstrair o I/O
type IOHandler interface {
	ReadInt() int
	WriteInt(value int)
	WriteStr(value string)
}
