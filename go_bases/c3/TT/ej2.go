/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el
main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/

package main

type usuarios struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []*producto
}
type producto struct {
	Nombre   string
	Tipo     string
	Cantidad int
}

func newProducto(nombre string, tipo string, cant int) *producto {
	p := new(producto)
	p.Nombre = nombre
	p.Tipo = tipo
	p.Cantidad = 0
	return p
}
func addProducto(u *usuarios, p *producto, cant int) {
	p.Cantidad = cant
	u.Productos = append(u.Productos, p)
}

func deleteProducto(u *usuarios) {
	for i := range u.Productos {
		u.Productos[i] = nil
	}
}
