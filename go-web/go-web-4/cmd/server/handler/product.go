package handler

import (
	"fmt"
	"net/http"
	"strconv"

	products "github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-web/go-web-4/internal/products/service"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "s3cr3t" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"result": "Invalid token",
			})
			return
		}
		pList, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, pList)
	}
}

func (c *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "s3cr3t" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"result": "Invalid token",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id",
			})
			return
		}
		p, err := c.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "s3cr3t" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"result": "Invalid token",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id",
			})
			return
		}
		var req request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if field, ok := CheckFields(&req); !ok {
			msg := fmt.Sprintf("Missing field %v", field)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": msg,
			})
			return
		}
		p, err := c.service.Update(id, req.Name, req.Color, req.Code, req.Price, req.Stock, req.Published)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "s3cr3t" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"result": "Invalid token",
			})
			return
		}
		var req request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if field, ok := CheckFields(&req); !ok {
			msg := fmt.Sprintf("Missing field %v", field)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": msg,
			})
			return
		}
		p, err := c.service.Store(req.Name, req.Color, req.Code, req.Price, req.Stock, req.Published)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "s3cr3t" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"result": "Invalid token",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id",
			})
			return
		}
		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	}
}

func (c *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "s3cr3t" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"result": "Invalid token",
			})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id",
			})
			return
		}
		var req request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.UpdateNameAndPrice(id, req.Name, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func CheckFields(req *request) (field string, ok bool) {
	if req.Name == "" {
		ok = false
		field = "Name"
		return
	} else if req.Color == "" {
		ok = false
		field = "Color"
		return
	} else if req.Price == 0 {
		ok = false
		field = "Price"
		return
	} else if req.Stock == 0 {
		ok = false
		field = "Stock"
		return
	} else if req.Code == "" {
		ok = false
		field = "Code"
		return
	} else if !req.Published {
		ok = false
		field = "Published"
		return
	}
	ok = true
	return
}
