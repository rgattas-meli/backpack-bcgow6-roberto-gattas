package main

import (
	"fmt"
)

func salario(min int, cat string) int {
	if min == 0 {
		return 0
	} else {
		if cat == "A" {
			return ((((min / 60) * 3000) * 3) / 2)
		} else if cat == "B" {
			return ((((min / 60) * 1500) * 12) / 10)
		} else if cat == "C" {
			return ((min / 60) * 1000)
		} else {
			return 0
		}
	}
}

func main() {
	fmt.Println(salario(60, "C"))
}
