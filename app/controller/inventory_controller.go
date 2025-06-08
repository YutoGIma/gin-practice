package controller

import (
	"myapp/app/model"
	"myapp/app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	inventoryUseCase usecase.InventoryUseCase
}

func NewInventoryController(inventoryUseCase usecase.InventoryUseCase) InventoryController {
	return InventoryController{
		inventoryUseCase: inventoryUseCase,
	}
}

func (c InventoryController) GetInventories(ctx *gin.Context) {
	inventories, err := c.inventoryUseCase.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, inventories)
}

// func (c InventoryController) GetInventoryDetail(ctx *gin.Context) {
// 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	inventory, err := c.inventoryUseCase.GetByID(uint(id))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, inventory)
// }

func (c InventoryController) CreateInventory(ctx *gin.Context) {
	var inventory model.Inventory
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdInventory, err := c.inventoryUseCase.Create(inventory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdInventory)
}

// func (c InventoryController) UpdateInventory(ctx *gin.Context) {
// 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	var inventory model.Inventory
// 	if err := ctx.ShouldBindJSON(&inventory); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	updatedInventory, err := c.inventoryUseCase.Update(uint(id), inventory)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, updatedInventory)
// }

func (c InventoryController) DeleteInventory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.inventoryUseCase.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Inventory deleted successfully"})
}
