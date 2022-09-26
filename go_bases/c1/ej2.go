package main

import "fmt"

var temperature float64
var humidity int
var pressure float64

func main() {
	temperature = 21.2
	humidity = 80
	pressure = 1
	fmt.Println(temperature, "°C")
	fmt.Println(humidity, "%")
	fmt.Println(pressure, "atm")
}
