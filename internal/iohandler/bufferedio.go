package iohandler

// BufferedIO é uma implementação de IOHandler bufferizada, para testes
type BufferedIO struct {
	inputInt  chan int
	outputInt chan int
	inputStr  chan string
	outputStr chan string
}

// ReadInt é implementação da interface IOHandler
func (io BufferedIO) ReadInt() int {
	return <-io.inputInt
}

// WriteInt é implementação da interface IOHandler
func (io BufferedIO) WriteInt(value int) {
	io.outputInt <- value
}

// WriteStr é implementação da interface IOHandler
func (io BufferedIO) WriteStr(value string) {
	io.outputStr <- value
}

// InputInt simula a entrada de um valor inteiro pelo usuário
func (io BufferedIO) InputInt(value int) {
	io.inputInt <- value
}

// OutputInt retorna uma saída do programa como inteiro
func (io BufferedIO) OutputInt() int {
	return <-io.outputInt
}

// OutputStr retorna uma saída do programa como string
func (io BufferedIO) OutputStr() string {
	return <-io.outputStr
}

// GetBufferedIO retorna uma nova instância de um IOHandler bufferizado
func GetBufferedIO() BufferedIO {
	return BufferedIO{
		inputInt:  make(chan int, 100),
		outputInt: make(chan int, 100),
		inputStr:  make(chan string, 100),
		outputStr: make(chan string, 100)}
}
