package products

import (
	"fmt"
	"time"

	store "github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-web/go-web-5/pkg/store"
)

type Product struct {
	Id           int       `json:"id"`
	Name         string    `json:"name" binding:"required"`
	Color        string    `json:"color"`
	Price        float64   `json:"price" binding:"required"`
	Stock        int       `json:"stock" binding:"required"`
	Code         string    `json:"code"`
	Published    bool      `json:"published"`
	CreationDate time.Time `json:"creation_date"`
}

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Store(id int, name, color, code string, price float64, stock int, published bool) (Product, error)
	GetLastId() (int, error)
	Update(id int, name, color, code string, price float64, stock int, published bool) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(id int, name, color, code string, price float64, stock int, published bool) (Product, error) {
	var pList []Product
	r.db.Read(&pList)
	p := Product{id, name, color, price, stock, code, published, time.Now()}
	pList = append(pList, p)
	if err := r.db.Write(pList); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var pList []Product
	r.db.Read(&pList)

	return pList, nil
}

func (r *repository) GetById(id int) (Product, error) {
	var pList []Product
	r.db.Read(&pList)
	for _, p := range pList {
		if p.Id == id {
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("Product not found")
}

func (r *repository) GetLastId() (int, error) {
	var pList []Product
	r.db.Read(&pList)
	if err := r.db.Read(&pList); err != nil {
		return 0, err
	}
	if len(pList) == 0 {
		return 0, nil
	}
	return pList[len(pList)-1].Id, nil
}

func (r *repository) Update(id int, name, color, code string, price float64, stock int, published bool) (Product, error) {
	var pList []Product
	r.db.Read(&pList)
	p, err := r.GetById(id)
	if err != nil {
		return Product{}, err
	}
	p.Name = name
	p.Color = color
	p.Code = code
	p.Price = price
	p.Stock = stock
	p.Published = published
	pList[id-1] = p
	if err := r.db.Write(pList); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	var pList []Product
	r.db.Read(&pList)
	p, err := r.GetById(id)
	if err != nil {
		return Product{}, err
	}
	p.Name = name
	p.Price = price
	pList[id-1] = p
	if err := r.db.Write(pList); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var pList []Product
	r.db.Read(&pList)
	p, err := r.GetById(id)
	if err != nil {
		return Product{}, err
	}
	p.Name = name
	pList[id-1] = p
	if err := r.db.Write(pList); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var pList []Product
	r.db.Read(&pList)
	_, err := r.GetById(id)
	if err != nil {
		return err
	}
	pList = append(pList[:id-1], pList[id:]...)
	fmt.Printf("%+v\n", pList)
	if err := r.db.Write(pList); err != nil {
		return err
	}
	return nil
}
