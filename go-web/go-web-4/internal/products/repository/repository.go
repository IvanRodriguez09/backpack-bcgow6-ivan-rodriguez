package products

import (
	"fmt"
	"time"
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

var pList []Product
var lastId int

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Store(id int, name, color, code string, price float64, stock int, published bool) (Product, error)
	GetLastId() (int, error)
	Update(id int, name, color, code string, price float64, stock int, published bool) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, name, color, code string, price float64, stock int, published bool) (Product, error) {
	p := Product{id, name, color, price, stock, code, published, time.Now()}
	pList = append(pList, p)
	lastId = p.Id
	return p, nil
}

func (r *repository) GetAll() ([]Product, error) {
	return pList, nil
}

func (r *repository) GetById(id int) (Product, error) {
	if id > len(pList) {
		return Product{}, fmt.Errorf("product %d not found", id)
	}
	for _, p := range pList {
		if p.Id == id {
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("product %d not found", id)
}

func (r *repository) GetLastId() (int, error) {
	if lastId == 0 {
		return lastId, nil
	} else {
		return lastId, nil
	}
}

func (r *repository) Update(id int, name, color, code string, price float64, stock int, published bool) (Product, error) {
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
	return p, nil
}

func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	p, err := r.GetById(id)
	if err != nil {
		return Product{}, err
	}
	p.Name = name
	p.Price = price
	pList[id-1] = p
	return p, nil
}

func (r *repository) Delete(id int) error {
	_, err := r.GetById(id)
	if err != nil {
		return err
	}
	pList = append(pList[:id-1], pList[id:]...)
	return nil
}
