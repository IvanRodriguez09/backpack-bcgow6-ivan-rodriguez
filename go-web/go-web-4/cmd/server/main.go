package main

import (
	"github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-web/go-web-4/cmd/server/handler"
	products "github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-web/go-web-4/internal/products/service"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)

	p := handler.NewProduct(service)

	router := gin.Default()
	pr := router.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetById())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())

}
