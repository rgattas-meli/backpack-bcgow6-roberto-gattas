package main

import (
	"errors"
	"fmt"
)

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func perros(cant int) float64 {
	return float64(cant * 10)
}
func gatos(cant int) float64 {
	return float64(cant * 5)
}
func hamsters(cant int) float64 {
	return float64(cant) * 0.25
}
func arañas(cant int) float64 {
	return float64(cant) * 0.15
}
func Animal(tipo string) (func(int) float64, error) {
	switch tipo {
	case dog:
		return perros, nil
	case cat:
		return gatos, nil
	case hamster:
		return hamsters, nil
	case spider:
		return arañas, nil
	default:
		return nil, errors.New("no such animal")
	}
}

func main() {
	animalDog, msg := Animal(dog)
	animalCat, msg := Animal(cat)
	if msg == nil {
		fmt.Println(animalDog(10))
		fmt.Println(animalCat(10))
	}
	fmt.Println(msg)

}
