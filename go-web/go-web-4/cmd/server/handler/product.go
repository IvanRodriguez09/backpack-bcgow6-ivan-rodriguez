package handler

import (
	"net/http"
	"strconv"

	products "github.com/IvanRodriguez09/backpack-bcgow6-ivan-rodriguez/go-web/go-web-4/internal/products/service"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name      string  `json:"name" binding:"required"`
	Color     string  `json:"color"`
	Price     float64 `json:"price" binding:"required"`
	Stock     int     `json:"stock" binding:"required"`
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
