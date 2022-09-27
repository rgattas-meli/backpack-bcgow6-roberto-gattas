/*
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
  - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
  - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/
package main

type producto struct {
	Nombre string
	Tipo   string
	Precio float64
}
type tienda struct {
	nombre string
	list   []producto
}
type Producto interface {
	CalcularCosto() int
}
type Ecommerce interface {
	Total() float64
	Agregar(producto)
}

func nuevoProducto(tipo, nombre string, precio float64) *producto {
	return &producto{
		Nombre: nombre,
		Tipo:   tipo,
		Precio: precio,
	}
}

func nuevaTienda(nombre string) *tienda {
	return &tienda{
		nombre: nombre,
		list:   []producto{},
	}
}

func (p *producto) CalcularCosto() float64 {
	var total float64
	total = 0
	if p.Tipo == "Pequeño" {
		total = p.Precio
	} else if p.Tipo == "Mediano" {
		total = p.Precio * 1.03
	} else {
		total = p.Precio*1.06 + 2500
	}
	return total
}
func (t *tienda) Total() float64 {
	total := 0.0
	for _, p := range t.list {
		total += p.CalcularCosto()
	}
	return total
}
func (t *tienda) Agregar(p producto) {
	t.list = append(t.list, p)
}

func main() {

}
