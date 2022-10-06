package products

/*Dentro del paquete deben estar las capas:
Servicio, debe contener la lógica de nuestra aplicación.
Se debe crear el archivo service.go.
Se debe generar la interface Service con todos sus métodos.
Se debe generar la estructura service que contenga el repositorio.
Se debe generar una función que devuelva el Servicio.
Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..).
*/

type Service interface {
	GetAll() ([]Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	Patch(id int, name string, price float64) (Product, error)
	Delete(id int) error

}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, tipo, cantidad, precio)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}


func (s *service) Update(id int, name, productType string, count int, price float64) (Product, error) {

	return s.repository.Update(id, name, productType, count, price)
 }
 
 func (s *service) Patch(id int, name string, price float64) (Product, error) {

	return s.repository.Patch(id, name, price)
 }
 

 func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
 }
 