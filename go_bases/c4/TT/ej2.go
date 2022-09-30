/* El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. 
Los datos requeridos para registrar a un cliente son:
Legajo
Nombre y Apellido
DNI
Número de teléfono
Domicilio

Tarea 1: El número de legajo debe ser asignado o generado por separado yen forma previa a la carga de los restantes gastos. 
Desarrolla e implementa una función para generar un ID que luego utilizarás  para asignarlo como valor a “Legajo”. 
Si por algún motivo esta función retorna valor “nil”, debe generar un panic que interrumpa la ejecución y aborte.
*/
package main
import(
	"fmt"
	"os"
	"errors"
)

type cliente struct{
	Legajo	int
	Nombre string
	Apellido  string
	DNI int
	telefono int
	domicilio string
}
//func newUser(nombre string, apellido  string)	{
func newUser(count int) string{
	count ++
	return string(count)

}

func openFile(archivo string)	{
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("el archivo indicado no fue encontrado o está dañado")
		}
		fmt.Println("lectura de archivo finalizada")

	}()
	_, err := os.Open(archivo)
	if err != nil {
		panic(err)
	}
}
func main() {
	openFile("customers.txt")
	count := 0
	id := newUser(count)
	if id == "nil"	{
		err := errors.New("error1")
		panic(err)
	}


}