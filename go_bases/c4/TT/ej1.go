package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("el archivo no fue encontrado o está dañado")

		}
		fmt.Println("ejecución finalizada")

	}()
	read, err := os.ReadFile("./customer.txt")
	if err != nil {
		panic("el archivo no fue encontrado o está dañado")
	}

	file := string(read)
	fmt.Println(file)
}
