/*
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados,
separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	records := [][]string{
		{"1011", ";", "100", ";", "15"},
		{"245642", ";", "2000", ";", "33"},
		{"18181814", ";", "345", ";", "3455"},
	}

	f, err := os.Create("productos.csv")
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records)
	fmt.Println("hola")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)

}
