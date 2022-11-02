package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

const name = "Ivan"

func readJSON() (products []Product) {

	jsonData, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(jsonData, &products); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {

	router := gin.Default()

	msg := fmt.Sprintf("Hello %s", name)

	router.GET("/helloGin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": msg,
		})
	})

	router.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": readJSON(),
		})
	})

	if err := router.Run(); err != nil {
		log.Fatal("Error: ", err)
	}
}
