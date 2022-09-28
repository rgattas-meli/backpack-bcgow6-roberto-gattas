/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que,
en reemplazo de “Error()”,  se implemente “errors.New()”.
*/
package main

import (
	"errors"
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

		fmt.Println(errors.New("La petición ha fallado."))
		return status

	} else {
		fmt.Printf("Status %d, Funciona!", status)
	}
	return status
}

var salary int

func main() {
	salary = 2100000
	myCustomErrorTest(salary) //llamamos a nuestra func
	//fmt.Printf("Status %d, Funciona!", status)

}
