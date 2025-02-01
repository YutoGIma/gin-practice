package controller

import (
	"myapp/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	InventoryService *service.InventoryService
}

func (ctrl *InventoryController) GetInventories(c *gin.Context) {
	inventories, err := ctrl.InventoryService.GetInventories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inventories)
}
