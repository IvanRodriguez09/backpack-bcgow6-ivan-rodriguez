package products

import product "github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-web/go-web-4/internal/products/repository"

type Service interface {
	GetAll() ([]product.Product, error)
	GetById(id int) (product.Product, error)
	Store(name, color, code string, price float64, stock int, published bool) (product.Product, error)
	GetLastId() (int, error)
	Update(id int, name, color, code string, price float64, stock int, published bool) (product.Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (product.Product, error)
	Delete(id int) error
}

type service struct {
	repository product.Repository
}

func NewService(r product.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]product.Product, error) {
	pList, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return pList, nil
}

func (s *service) GetById(id int) (product.Product, error) {
	p, err := s.repository.GetById(id)
	if err != nil {
		return product.Product{}, err
	}
	return p, nil
}

func (s *service) Store(name, color, code string, price float64, stock int, published bool) (product.Product, error) {
	lastId, err := s.repository.GetLastId()
	if err != nil {
		return product.Product{}, err
	}
	lastId++
	prod, err := s.repository.Store(lastId, name, color, code, price, stock, published)
	if err != nil {
		return product.Product{}, err
	}
	return prod, nil
}

func (s *service) GetLastId() (int, error) {
	return s.repository.GetLastId()
}

func (s *service) Update(id int, name, color, code string, price float64, stock int, published bool) (product.Product, error) {
	return s.repository.Update(id, name, color, code, price, stock, published)
}

func (s *service) UpdateNameAndPrice(id int, name string, price float64) (product.Product, error) {
	return s.repository.UpdateNameAndPrice(id, name, price)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
