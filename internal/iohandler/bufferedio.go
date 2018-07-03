package iohandler

type bufferedIO struct {
	inputInt  chan int
	outputInt chan int
}

func (io bufferedIO) ReadInt() int {
	return <-io.inputInt
}

func (io bufferedIO) WriteInt(value int) {
	io.outputInt <- value
}

// GetBufferedIO retorna uma nova instância de um IOHandler bufferizado
func GetBufferedIO() IOHandler {
	return bufferedIO{
		inputInt:  make(chan int, 100),
		outputInt: make(chan int, 100)}
}
