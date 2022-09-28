/*
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error:
el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000.
Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/
package main

import (
	"fmt"
	"os"
)

type myCustomError struct {
	status int
	msg    string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func myCustomErrorTest(status int) (int, error) {
	if status < 150000 {
		return 404, &myCustomError{
			status: status,
			msg:    "algo salió mal",
		}
	}
	return 20000000, nil
}

var salary int

func main() {
	salary = 2100000
	status, err := myCustomErrorTest(salary) //llamamos a nuestra func
	if err != nil {                          //hacemos una validación del valor de err
		fmt.Println(err) //si err no es nil, imprimimos el error y...
		os.Exit(1)       //utilizamos este método para salir del programa
	}
	fmt.Printf("Status %d, Funciona!", status)

}
