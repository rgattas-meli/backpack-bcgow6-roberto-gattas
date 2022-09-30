package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("el archivo indicado no fue encontrado o está dañado")

		}
		fmt.Println("ejecución finalizada")

	}()
	_, err := os.Open("customers.txt")
	if err != nil {
		panic(err)
	}

}
