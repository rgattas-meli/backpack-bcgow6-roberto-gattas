package main

import (
	"errors"
	"fmt"
)

func promedio(alumnos ...int) (int, error) {
	if len(alumnos) == 0 {
		return 0, errors.New("No alumnos")
	} else {
		count := 0
		for i := 0; i < len(alumnos); i++ {
			if alumnos[i] < 0 {
				return 0, errors.New("Numero negativo")
			} else {
				count += alumnos[i]
			}
		}
		count = count / len(alumnos)
		return count, nil
	}
}

func main() {
	count, err := promedio(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}
}
