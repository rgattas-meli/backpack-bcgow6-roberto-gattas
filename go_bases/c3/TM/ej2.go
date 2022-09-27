package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const PATH = "productos.csv"

func main() {
	f, err := os.Open(PATH)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(f)
	record, _ := reader.ReadAll()
	total := 0

	for _, prod := range record {
		precio, _ := strconv.Atoi(prod[2])
		cant, _ := strconv.Atoi(prod[4])
		total += precio * cant
		fmt.Println(prod[0] + "\t" + prod[2] + "\t" + prod[4] + "\t")
	}
	fmt.Println("El total es:" + strconv.Itoa(total))

}
