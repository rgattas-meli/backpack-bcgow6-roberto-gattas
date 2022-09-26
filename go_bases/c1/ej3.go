package main

import "fmt"

/*
var 1nombre string  -- esta mal debido a tener 1 como primer char en nombre de variable
var apellido string
var int edad		-- esta mal debido mal formateo de definicion
1apellido := 6    -- esta mal debido a numero como 1er caracter de nombre variable , fuera de main no es correcto
var licencia_de_conducir = true
var estatura de la persona int -- esta mal debido a espacios en nombre de variable
cantidadDeHijos := 2  -- fuera de main es incorrecto
*/

var nombre string
var apellido string
var edad int
var licencia_de_conducir = true
var estatura_de_la_persona int

func main() {
	apellido2 := 6
	cantidadDeHijos := 2
	fmt.Println(nombre, apellido, apellido2, edad, licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)
}
