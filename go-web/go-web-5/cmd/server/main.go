package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-web/go-web-5/cmd/server/handler"
	repository "github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-web/go-web-5/internal/products/repository"
	service "github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-web/go-web-5/internal/products/service"
	store "github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-web/go-web-5/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	fmt.Println("holaa ", os.Getenv("TOKEN"))
	db := store.New(store.FileType, "./products.json")
	repo := repository.NewRepository(db)
	service := service.NewService(repo)

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
