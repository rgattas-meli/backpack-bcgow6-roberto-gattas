package main

import "fmt"

type matrix struct {
	arr        [][]float64
	alto       int
	ancho      int
	cuadratica bool
	vmax       float64
}

func (m *matrix) Set(valores []float64) {
	m.arr = append(m.arr, valores)
	m.alto++

}
func (m *matrix) Print() {
	for p := 0; p < m.alto; p++ {
		for i := 0; i < m.ancho; i++ {
			fmt.Print(m.arr[p][i], " ")
		}
		fmt.Print("\n")
	}
}

func main() {

}
