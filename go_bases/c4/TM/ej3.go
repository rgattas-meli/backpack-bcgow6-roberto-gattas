/*
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”,
para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza
el mínimo imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible
es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/
package main

import (
	"fmt"
)

type myCustomError struct {
	status int
	msg    string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func myCustomErrorTest(status int) int {
	if status < 150000 {

		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", status)
		fmt.Println(" Error ocurrido: ", err)
		return status

	} else {
		fmt.Printf("Funciona! y tiene valor %d", status)
	}
	return status
}

var salary int

func main() {
	salary = 2100
	myCustomErrorTest(salary) //llamamos a nuestra func
	//fmt.Printf("Status %d, Funciona!", status)

}
