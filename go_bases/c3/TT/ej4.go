package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(array []int) []int {
	now := time.Now()
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println("Termina el programa")
	fmt.Printf("El progrma 1 demoro: %d ms\n", time.Now().Sub(now).Milliseconds())
	return array
}
func Selection_Sort(array []int) []int {
	now := time.Now()
	var min_index int
	var temp int
	size := len(array)
	for i := 0; i < size-1; i++ {
		min_index = i
		// Find index of minimum element
		for j := i + 1; j < size; j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		temp = array[i]
		array[i] = array[min_index]
		array[min_index] = temp
	}
	fmt.Println("Termina el programa")
	fmt.Printf("El progrma 2 demoro: %d ms\n", time.Now().Sub(now).Milliseconds())
	return array
}
func insertionsort(items []int) {
	now := time.Now()
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	fmt.Println("Termina el programa")
	fmt.Printf("El progrma  3 demoro: %d ms\n", time.Now().Sub(now).Milliseconds())
}

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)
	time.Now()
	for i := 0; i < 3; i++ {
		if i == 0 {
			insertionsort(variable1)
			BubbleSort(variable1)
			Selection_Sort(variable1)
		} else if i == 1 {
			insertionsort(variable2)
			BubbleSort(variable2)
			Selection_Sort(variable2)

		} else {
			insertionsort(variable3)
			BubbleSort(variable3)
			Selection_Sort(variable3)
		}
	}

}
