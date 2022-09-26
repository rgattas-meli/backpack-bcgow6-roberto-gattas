package main

import "fmt"

func impuestos(salario float64) float64 {

	if salario > 50000 && salario < 150001 {
		return salario * 0.17
	} else if salario > 150000 {
		return salario * 0.27
	} else {
		return 0
	}
}

func main() {

	var salario float64
	salario = 80000
	fmt.Println(impuestos(salario))
}
