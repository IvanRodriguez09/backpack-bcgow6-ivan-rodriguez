package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-testing/go-testing-3/cmd/server/handler"
	products "github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-testing/go-testing-3/internal/products"
	store "github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-testing/go-testing-3/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	fmt.Println("holaa ", os.Getenv("TOKEN"))
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)

	p := handler.NewProduct(service)

	router := gin.Default()
	pr := router.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetById())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNameAndPrice())
	pr.DELETE("/:id", p.Delete())

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
