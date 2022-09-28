/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos,
Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que
el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

*/

package main

import "fmt"

type producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}
type servicio struct {
	Nombre  string
	Precio  float64
	minutos int
}
type mantenimiento struct {
	Nombre string
	Precio float64
}

func addProducto(productos []producto, c chan float64) {
	total := 0.0
	for i := 0; i < len(productos); i++ {
		total += (productos[i].Precio) * float64(productos[i].Cantidad)
	}
	c <- total
}

func addServicio(servicios []servicio, c chan float64) {
	total := 0.0
	for i := 0; i < len(servicios); i++ {
		if servicios[i].minutos < 30 {
			total += (servicios[i].Precio)
		} else {
			total += (servicios[i].Precio) * float64(servicios[i].minutos)
			total /= 30.0
		}
	}
	c <- total

}

func addMantenimiento(mantenimientos []mantenimiento, c chan float64) {
	total := 0.0
	for i := 0; i < len(mantenimientos); i++ {
		total += (mantenimientos[i].Precio)
	}
	c <- total
}

func main() {
	var m []mantenimiento
	var s []servicio
	var p []producto
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go addProducto(p, c1)
	go addServicio(s, c2)
	go addMantenimiento(m, c3)
	total := <-c1 + <-c2 + <-c3
	fmt.Println("El total es", total)

}
