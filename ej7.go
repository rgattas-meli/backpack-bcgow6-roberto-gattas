package main
import "fmt"

var meses = []string{"enero","febrero","marzo","abril","mayo","junio","julio", "agosto","septiembre","octubre","noviembre","diciembre"}
var numero int
func main() {
	numero = 4
	if ((numero > 0) && (numero < 13)) {
		fmt.Println(numero, ",", meses[numero-1])
	} else {
		fmt.Println("Numero inválido")
	}
}


//hay otra solucion posible, la cual seria utilizando switch, esta no la usaria debido a la ineficiencia de escribir cada
// mes y a su vez cada impresion debido a cada mes, a su vez se podria usar slice, el cual seria tambien eficiente y rapido