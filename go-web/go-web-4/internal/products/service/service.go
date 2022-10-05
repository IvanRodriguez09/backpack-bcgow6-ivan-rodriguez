package products

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Store(name, color, code string, price float64, stock int, published bool) (Product, error)
	GetLastId() (int, error)
	Update(id int, name, color, code string, price float64, stock int, published bool) (Product, error)
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
	pList, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return pList, nil
}

func (s *service) GetById(id int) (Product, error) {
	p, err := s.repository.GetById(id)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (s *service) Store(id int, name, color, code string, price float64, stock int, published bool) (Product, error) {
	lastId, err := s.repository.GetLastId()
	if err != nil {
		return Product{}, err
	}
	lastId++
	product, err := s.repository.Store(lastId, name, color, code, price, stock, published)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) GetLastId() (int, error) {
	return s.repository.LastId()
}

func (s *service) Update(id int, name, color, code string, price float64, stock int, published bool) (Product, error) {
	return s.repository.Update(id, name, color, code, price, stock, published)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
