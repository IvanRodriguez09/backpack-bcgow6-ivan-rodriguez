package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Color        string    `json:"color"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	Code         string    `json:"code"`
	Published    bool      `json:"published"`
	CreationDate time.Time `json:"creation_date"`
}

func readJSON() (products []Product) {

	jsonData, err := os.ReadFile("../go-web-1/products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(jsonData, &products); err != nil {
		log.Fatal(err)
	}

	return
}

var products []Product = readJSON()

func main() {

	router := gin.Default()
	router.GET("/", HandlerHome)
	productsRoutes := router.Group("/products")
	{
		productsRoutes.GET("/:id", HandlerGetById)
		productsRoutes.GET("/filter1", HandlerGetProductByPriceAndColor)
	}
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}

func HandlerHome(c *gin.Context) {
	c.String(http.StatusOK, "Server Ok!")
}

func HandlerGetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id >= len(products) {
		ctx.String(400, "Bad request by id; %s\n", ctx.Param("id"))
	} else {
		product := products[id]
		if product.Id == 0 {
			ctx.String(404, "Product not found by id %s: \n", ctx.Param("id"))
		} else {
			ctx.String(200, "Product found by id %s, details: %+v\n", ctx.Param("id"), product)
		}
	}
}

func HandlerGetProductByPriceAndColor(ctx *gin.Context) {
	var resFilter []Product

	color := ctx.Query("color")
	price, err := strconv.ParseFloat(ctx.Query("price"), 64)

	if err != nil {
		ctx.String(400, "Bad Request by price: %v", ctx.Query("price"))
	}
	for _, product := range products {
		if product.Color == color && product.Price == price {
			resFilter = append(resFilter, product)
		}
	}
	ctx.String(200, "Product list found: %+v", resFilter)

}

func QueryTest(ctx *gin.Context) {
	isActive := ctx.Query("active")
	ctx.JSON(http.StatusOK, gin.H{"status": isActive})
}
