package controller

import (
	"myapp/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	InventoryService *service.InventoryService
}

func (c *InventoryController) GetInventories(ctx *gin.Context) {
	inventories, err := c.InventoryService.GetInventories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, inventories)
}
