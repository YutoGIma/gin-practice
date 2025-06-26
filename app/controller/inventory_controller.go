package controller

import (
	"myapp/app/model"
	"myapp/app/usecase"
	"myapp/app/usecase/request"
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

// GetInventories godoc
// @Summary 在庫一覧取得
// @Description すべての在庫情報を取得します
// @Tags inventories
// @Accept json
// @Produce json
// @Success 200 {array} model.Inventory
// @Failure 500 {object} map[string]string
// @Router /inventories [get]
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

// CreateInventory godoc
// @Summary 在庫作成
// @Description 新しい在庫レコードを作成します
// @Tags inventories
// @Accept json
// @Produce json
// @Param inventory body model.Inventory true "在庫情報"
// @Success 201 {object} model.Inventory
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /inventories [post]
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

// UpdateInventory godoc
// @Summary 在庫更新
// @Description 指定したIDの在庫情報を更新します
// @Tags inventories
// @Accept json
// @Produce json
// @Param id path int true "在庫ID"
// @Param request body request.InventoryUpdateRequest true "在庫更新情報"
// @Success 200 {object} model.Inventory
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /inventories/{id} [put]
func (c InventoryController) UpdateInventory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	var req request.InventoryUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inventory, err := c.inventoryUseCase.Update(uint(id), req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, inventory)
}

func (c InventoryController) UpdateInventoryOnPurchase(ctx *gin.Context) {
	var req request.InventoryPurchaseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.inventoryUseCase.UpdateOnPurchase(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
}

// DeleteInventory godoc
// @Summary 在庫削除
// @Description 指定したIDの在庫を削除します
// @Tags inventories
// @Accept json
// @Produce json
// @Param id path int true "在庫ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /inventories/{id} [delete]
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

// RestockInventory godoc
// @Summary 商品入荷
// @Description 指定した商品の在庫を追加します
// @Tags inventories
// @Accept json
// @Produce json
// @Param request body request.InventoryRestockRequest true "入荷情報"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /inventories/restock [post]
func (c InventoryController) RestockInventory(ctx *gin.Context) {
	var req request.InventoryRestockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inventory, err := c.inventoryUseCase.RestockInventory(req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "在庫を入荷しました",
		"inventory": inventory,
	})
}
