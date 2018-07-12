package obi2017

import (
	"github.com/ibraimgm/goobi/internal/iohandler"
	"github.com/ibraimgm/goobi/internal/obiinfo"
)

// Castelos é a solução para o exercício 'castelos' da OBI 2017/2
var Castelos = obiinfo.New("castelos", 2017, 2, "https://olimpiada.ic.unicamp.br/pratique/p1/2017/f2/castelos/", castelos)

type intSet map[int]bool

func newIntSet() intSet {
	return make(map[int]bool)
}

func (set intSet) put(value int) {
	set[value] = true
}

func (set intSet) has(value int) bool {
	return set[value]
}

func (set intSet) clone() intSet {
	other := intSet{}

	for k, v := range set {
		other[k] = v
	}

	return other
}

type estrada struct {
	de   intSet
	para intSet
}

func castelos(handler iohandler.IOHandler) {
	// numero de castelos e numero de ordens
	n := handler.ReadInt()
	m := handler.ReadInt()

	// inicializa os caminhos
	caminhos := map[int]estrada{}
	for i := 0; i < n; i++ {
		caminhos[i+1] = estrada{newIntSet(), newIntSet()}
	}

	// caminhos entre os castelos
	for i := 0; i < n-1; i++ {
		u := handler.ReadInt()
		v := handler.ReadInt()

		caminhos[u].para.put(v)
		caminhos[v].de.put(u)
	}

	// pintura inicial
	pintura := make([]int, n)

	for i := 0; i < m; i++ {
		p := handler.ReadInt()
		q := handler.ReadInt()
		c := handler.ReadInt()

		solution := walk(p, q, caminhos, intSet{})
		paint(pintura, solution, c)
	}

	for _, v := range pintura {
		handler.WriteInt(v)
	}
}

func walk(from, to int, caminhos map[int]estrada, solution intSet) intSet {
	if from == to {
		solution.put(to)
		return solution
	}

	mySolution := solution.clone()
	mySolution.put(from)

	caminho := caminhos[from]

	if finalSolution := walkInPath(to, caminhos, caminho.de, mySolution.clone()); finalSolution != nil {
		return finalSolution
	}

	return walkInPath(to, caminhos, caminho.para, mySolution.clone())
}

func walkInPath(destino int, caminhos map[int]estrada, caminho, solution intSet) intSet {
	for origem, existe := range caminho {
		if existe && !solution.has(origem) {
			currentSolution := walk(origem, destino, caminhos, solution.clone())

			if currentSolution != nil {
				return currentSolution
			}
		}
	}

	return nil
}

func paint(pintura []int, set intSet, color int) {
	for k, v := range set {
		if v {
			pintura[k-1] = color
		}
	}
}
