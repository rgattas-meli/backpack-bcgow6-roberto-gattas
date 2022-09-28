/*
Una empresa de redes sociales requiere implementar una estructura 
usuario con funciones que vayan agregando información a la estructura. 
Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en 
memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.
*/
package main

import (
    "fmt"
    "os"

)
type usuarios struct {
	Nombre string
    Apellido string
    Edad int
    Correo string
    Contraseña string
}
func setNombre (u *usuarios, nombre2, apellido2 string)	{
	if u !=	 nil	{
		u.Nombre = nombre2
    	u.Apellido = apellido2
	}
}
func setEdad (u *usuarios, edad2 int)    {
	if u !=     nil    {
        u.Edad = edad2
    }
}
func setCorreo (u *usuarios, correo2 string) {
	if u!=     nil    {
        u.Correo = correo2
    }
}
func setContraseña (u *usuarios, contraseña string)	{
	if u!=     nil    {
        u.Contraseña = contraseña
    }
}
func main () {


}