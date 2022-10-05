package products

/*
Se debe crear el archivo repository.go
Se debe crear la estructura de la entidad
Se deben crear las variables globales donde guardar las entidades
Se debe generar la interface Repository con todos sus métodos
Se debe generar la estructura repository
Se debe generar una función que devuelva el Repositorio
Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
*/
import "fmt"
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var ps []Product
var LastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	LastID() (int, error)
	Patch(id int, name string, price float64) (Product, error)
	Update(id int,  name, productType string, count int, price float64) (Product, error)
	Delete(id int) error


}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	p := Product{id, nombre, tipo, cantidad, precio}
	ps = append(ps, p)
	LastID = id
	return p, nil

}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return LastID, nil
}

func (r *repository) Update(id int,  name, productType string, count int, price float64) (Product, error) {
	p := Product{Name: name, Type: productType, Count: count,Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}
	return p, nil
 }

 
 func (r *repository) Patch(id int, name string, price float64) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			ps[i].Price = price
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}
	return p, nil
 }
 
 func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil
 }

 