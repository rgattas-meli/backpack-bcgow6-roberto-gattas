package main

import (
	"fmt"
	"time"
)

type estudiante struct {
	nombre, apellido string
	dni              int
	fecha_ingreso    time.Time
}

func (e *estudiante) detalle(nombre, apellido string, dni int, fecha time.Time) {
	e.nombre = nombre
	e.apellido = apellido
	e.dni = dni
	e.fecha_ingreso = fecha
}

func main() {
	var es estudiante
	tiempo := time.Now()
	es.detalle("hola", "como", 20, tiempo)
	fmt.Println(es.nombre, es.apellido, es.dni, es.fecha_ingreso)
}
