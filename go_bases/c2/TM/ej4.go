package main

import (
	"errors"
	"fmt"
)

const (
	MINIMUM = "minimum"
	average = "average"
	maximum = "maximum"
)

func mini(alumnos ...int) int {
	if len(alumnos) < 0 {
		return -1
	} else {
		min := 40
		for i := 0; i < len(alumnos); i++ {
			if alumnos[i] < min {
				min = alumnos[i]
			}
		}
		return min
	}
}
func averg(alumnos ...int) int {
	avg := 0
	for i := 0; i < len(alumnos); i++ {
		avg += alumnos[i]
	}
	avg /= len(alumnos)
	return avg

}
func maxi(alumnos ...int) int {
	max := 0
	for i := 0; i < len(alumnos); i++ {
		if alumnos[i] > max {
			max = alumnos[i]
		}
	}
	return max
}

func operation(oper string) (func(...int) int, error) {
	if oper == MINIMUM {
		return mini, nil
	} else if oper == average {
		return averg, nil
	} else if oper == maximum {
		return maxi, nil
	} else {
		return nil, errors.New("Tipo de calculo incorrecto")
	}
}

func main() {
	minFunc, err := operation(MINIMUM)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	if err == nil {
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println(minValue)
		fmt.Println(averageValue)
		fmt.Println(maxValue)
	}
}
