package controller

import (
	"myapp/app/model"
	"myapp/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	InventoryService service.InventoryService
	ProductService   service.ProductService
}

func (c InventoryController) GetInventories(ctx *gin.Context) {
	inventories, err := c.InventoryService.GetInventories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, inventories)
}

func (c InventoryController) CreateInventory(ctx *gin.Context) {
	var inventory model.Inventory
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := c.ProductService.GetProductDetail(inventory.ProductId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	inventory.Product = product
	if err := c.InventoryService.CreateInventory(inventory); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, inventory)
}

func (c InventoryController) UpdateInventory(ctx *gin.Context) {
	var inventory model.Inventory
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		return
	}
	inventory.ID = uint(id)
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := c.ProductService.GetProductDetail(inventory.ProductId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	inventory.Product = product
	err = c.InventoryService.UpdateInventory(inventory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, inventory)
}

func (c InventoryController) DeleteInventory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		return
	}
	err = c.InventoryService.DeleteInventory(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Inventory deleted"})
}
