package products

import "github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/tree/main/go_storage/c2/internal/domains"

type Service interface {
	Store(p domains.Product) (domains.Product, error)

	GetAll() ([]domains.Product, error)
	Delete(id int) error
}

type service struct {
	product Repository
}

func NewService(product Repository) Service {
	return &service{
		product: product,
	}
}

func (s *service) Store(p domains.Product) (domains.Product, error) {
	product, err := s.product.Store(p)
	if err != nil {
		return domains.Product{}, err
	}

	p.ID = product.ID
	return p, nil
}



func (s *service) GetAll() ([]domains.Product, error) {
	product, err := s.product.GetAll()
	if err != nil {
		return []domains.Product{}, err
	}

	return product, err
}


func (s *service) Delete(id int) error {
	err := s.product.Delete(id)
	if err != nil {
		return err
	}
	return nil
}