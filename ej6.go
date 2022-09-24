package main
import "fmt"

var (
	edad int
	trabaja bool
	antiguedad float64
	salario int

)
func main() {
	edad = 23
    trabaja = true
    antiguedad = 0.5
    salario = 1000000
	if ((edad > 22) && (trabaja) && (antiguedad >= 1))	{
		if (salario > 100000){
			fmt.Println("Prestamo otorgado sin intereses")
		} else {
			fmt.Println("Prestamo otorgado con intereses")
		}
	} else {
		fmt.Println("Prestamo no otorgado")
	}

}