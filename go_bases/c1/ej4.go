package main

import "fmt"

/*
var apellido string = "Gomez"
var edad int = "35"	-- esta mal debido a uso de string en declaracion de int
boolean := "false";  -- esta mal debido a ; y uso de nombre reservado, adentro de main es correcto, fuera falta declaracion var
var sueldo string = 45857.90 -- incorrecto uso de string , mediante uso de float en la declaracion
var nombre string = "Julián"
*/
var apellido string = "Gomez"
var edad int = 35
var sueldo string = "45857.90"
var nombre string = "Julián"

func main() {
	booleano := "false"

	fmt.Println(apellido)
	fmt.Println(edad)
	fmt.Println(sueldo)
	fmt.Println(nombre)
	fmt.Println(booleano)
}
