package obi2017

import (
	"fmt"
	"testing"

	"github.com/ibraimgm/goobi/internal/iohandler"
)

func TestCastelos_Ex1(t *testing.T) {
	caminhos := []int{
		1, 2,
		2, 5,
		2, 4,
		4, 3}

	ordens := []int{
		5, 3, 1,
		2, 4, 3}

	castelosTestHelper(t, 5, 2, caminhos, ordens, []int{0, 3, 1, 3, 1})
}

func TestCastelos_Ex2(t *testing.T) {
	caminhos := []int{
		7, 1,
		1, 2,
		6, 2,
		9, 6,
		2, 5,
		8, 6,
		4, 5,
		3, 4}

	ordens := []int{
		1, 4, 2,
		7, 8, 4,
		3, 4, 1,
		9, 2, 8}

	castelosTestHelper(t, 9, 4, caminhos, ordens, []int{4, 8, 1, 1, 2, 8, 4, 4, 8})
}

func castelosTestHelper(t *testing.T, n, m int, caminhos, ordens, esperado []int) {
	handler := iohandler.GetBufferedIO()

	handler.InputInt(n)
	handler.InputInt(m)

	for _, v := range caminhos {
		handler.InputInt(v)
	}

	for _, v := range ordens {
		handler.InputInt(v)
	}

	castelos(handler)

	for i, v := range esperado {
		x := handler.OutputInt()
		fmt.Println(x)

		if x != v {
			t.Errorf("Erro na posição %d. Esperado %d, encontrado %d.\n", i, v, x)
		}
	}
}
