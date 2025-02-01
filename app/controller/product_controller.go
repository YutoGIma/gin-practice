package controller

import (
	"myapp/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *service.ProductService
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
	products, err := c.ProductService.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
