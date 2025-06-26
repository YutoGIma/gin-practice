package controller

import (
	"myapp/app/usecase"
	"myapp/app/usecase/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PriceSettingController struct {
	priceSettingUseCase usecase.PriceSettingUseCase
}

func NewPriceSettingController(priceSettingUseCase usecase.PriceSettingUseCase) PriceSettingController {
	return PriceSettingController{
		priceSettingUseCase: priceSettingUseCase,
	}
}

// CreatePriceSetting godoc
// @Summary 価格設定作成
// @Description 指定された在庫に対して価格設定を作成します
// @Tags price-settings
// @Accept json
// @Produce json
// @Param id path int true "在庫ID"
// @Param request body request.CreatePriceSettingRequest true "価格設定作成リクエスト"
// @Success 201 {object} model.PriceSetting
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /inventories/{id}/prices [post]
func (c PriceSettingController) CreatePriceSetting(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDフォーマットです"})
		return
	}

	var req request.CreatePriceSettingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.InventoryID = uint(id)

	priceSetting, err := c.priceSettingUseCase.CreatePriceSetting(req)
	if err != nil {
		switch err.Error() {
		case "在庫が見つかりません":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusCreated, priceSetting)
}

// GetPriceSettingsByInventoryID godoc
// @Summary 在庫の価格設定履歴取得
// @Description 指定された在庫のすべての価格設定履歴を取得します
// @Tags price-settings
// @Accept json
// @Produce json
// @Param id path int true "在庫ID"
// @Success 200 {array} model.PriceSetting
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /inventories/{id}/prices [get]
func (c PriceSettingController) GetPriceSettingsByInventoryID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDフォーマットです"})
		return
	}

	priceSettings, err := c.priceSettingUseCase.GetPriceSettingsByInventoryID(uint(id))
	if err != nil {
		switch err.Error() {
		case "在庫が見つかりません":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, priceSettings)
}

// GetCurrentPriceSetting godoc
// @Summary 現在の価格設定取得
// @Description 指定された在庫の現在有効な価格設定を取得します
// @Tags price-settings
// @Accept json
// @Produce json
// @Param id path int true "在庫ID"
// @Success 200 {object} model.PriceSetting
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /inventories/{id}/prices/current [get]
func (c PriceSettingController) GetCurrentPriceSetting(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDフォーマットです"})
		return
	}

	priceSetting, err := c.priceSettingUseCase.GetCurrentPriceSetting(uint(id))
	if err != nil {
		switch err.Error() {
		case "在庫が見つかりません", "有効な価格設定が見つかりません":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, priceSetting)
}

// UpdatePriceSetting godoc
// @Summary 価格設定更新
// @Description 指定された価格設定を更新します
// @Tags price-settings
// @Accept json
// @Produce json
// @Param id path int true "在庫ID"
// @Param price_id path int true "価格設定ID"
// @Param request body request.UpdatePriceSettingRequest true "価格設定更新リクエスト"
// @Success 200 {object} model.PriceSetting
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /inventories/{id}/prices/{price_id} [put]
func (c PriceSettingController) UpdatePriceSetting(ctx *gin.Context) {
	priceID, err := strconv.ParseUint(ctx.Param("price_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効な価格設定IDフォーマットです"})
		return
	}

	var req request.UpdatePriceSettingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	priceSetting, err := c.priceSettingUseCase.UpdatePriceSetting(uint(priceID), req)
	if err != nil {
		switch err.Error() {
		case "価格設定が見つかりません":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, priceSetting)
}

// DeletePriceSetting godoc
// @Summary 価格設定削除
// @Description 指定された価格設定を削除します
// @Tags price-settings
// @Accept json
// @Produce json
// @Param id path int true "在庫ID"
// @Param price_id path int true "価格設定ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /inventories/{id}/prices/{price_id} [delete]
func (c PriceSettingController) DeletePriceSetting(ctx *gin.Context) {
	priceID, err := strconv.ParseUint(ctx.Param("price_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効な価格設定IDフォーマットです"})
		return
	}

	err = c.priceSettingUseCase.DeletePriceSetting(uint(priceID))
	if err != nil {
		switch err.Error() {
		case "価格設定が見つかりません":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}
