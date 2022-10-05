package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func createProductArray() (prods []Product) {
	prods = []Product{
		{
			Id:           1,
			Name:         "Iphone 13",
			Color:        "Black",
			Price:        3000,
			Stock:        10,
			Code:         "12abc",
			Published:    true,
			CreationDate: time.Now(),
		},
		{
			Id:           2,
			Name:         "Iphone 12",
			Color:        "Black",
			Price:        3000,
			Stock:        12,
			Code:         "12dfg",
			Published:    true,
			CreationDate: time.Now(),
		},
		{
			Id:           3,
			Name:         "Iphone 11",
			Color:        "Black",
			Price:        2000,
			Stock:        5,
			Code:         "12jkl",
			Published:    false,
			CreationDate: time.Now(),
		},
	}

	return
}

var products []Product = createProductArray()

func main() {

	router := gin.Default()
	router.POST("/products", HandlerCreateProduct)
	router.GET("/products", HandlerGetProducts)
	router.Run(":8080")
}

func HandlerCreateProduct(ctx *gin.Context) {
	var productJson *Product
	if ok := CheckToken(ctx.GetHeader("token")); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "no tiene permisos para realizar la petición solicitada",
		})
		return
	}
	if err := ctx.ShouldBindJSON(&productJson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ok, field := CheckFields(productJson); !ok {
		msg := fmt.Sprintf("El campo %s es requerido", field)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	lastId := products[len(products)-1].Id
	newProduct := Product{
		Id:           lastId + 1,
		Name:         productJson.Name,
		Color:        productJson.Color,
		Price:        productJson.Price,
		Stock:        productJson.Stock,
		Code:         productJson.Code,
		Published:    productJson.Published,
		CreationDate: time.Now(),
	}
	products = append(products, newProduct)
	ctx.JSON(http.StatusOK, newProduct)
}

func CheckFields(product *Product) (ok bool, field string) {
	if product.Name == "" {
		ok = false
		field = "Name"
		return
	} else if product.Color == "" {
		ok = false
		field = "Color"
		return
	} else if product.Price == 0 {
		ok = false
		field = "Price"
		return
	} else if product.Stock == 0 {
		ok = false
		field = "Stock"
		return
	} else if product.Code == "" {
		ok = false
		field = "Code"
		return
	} else if !product.Published {
		ok = false
		field = "Published"
		return
	}
	ok = true
	return
}

func CheckToken(token string) (ok bool) {
	if token == "s3cr3t" {
		ok = true
		return
	}
	return
}

func HandlerGetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
