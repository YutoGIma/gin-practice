package controller

import (
	"myapp/app/usecase"
	"myapp/app/usecase/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderController(orderUseCase usecase.OrderUseCase) OrderController {
	return OrderController{orderUseCase: orderUseCase}
}

// GetOrders godoc
// @Summary 注文一覧取得
// @Description すべての注文情報を取得します
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} model.Order
// @Failure 500 {object} map[string]string
// @Router /orders [get]
func (c *OrderController) GetOrders(ctx *gin.Context) {
	orders, err := c.orderUseCase.GetOrders()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

// GetOrderDetail godoc
// @Summary 注文詳細取得
// @Description 指定したIDの注文情報を取得します
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "注文ID"
// @Success 200 {object} model.Order
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /orders/{id} [get]
func (c *OrderController) GetOrderDetail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	order, err := c.orderUseCase.GetOrderByID(uint(id))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, order)
}

// GetUserOrders godoc
// @Summary ユーザーの注文一覧取得
// @Description 指定したユーザーの注文一覧を取得します
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "ユーザーID"
// @Success 200 {array} model.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id}/orders [get]
func (c *OrderController) GetUserOrders(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なユーザーIDです"})
		return
	}

	orders, err := c.orderUseCase.GetOrdersByUserID(uint(userID))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

// CreateOrder godoc
// @Summary 注文作成
// @Description 新しい注文を作成し、在庫を減らします
// @Tags orders
// @Accept json
// @Produce json
// @Param order body request.CreateOrderRequest true "注文情報"
// @Success 201 {object} model.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders [post]
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var req request.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := c.orderUseCase.CreateOrder(req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, order)
}

// CancelOrder godoc
// @Summary 注文キャンセル
// @Description 指定したIDの注文をキャンセルします
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "注文ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id}/cancel [post]
func (c *OrderController) CancelOrder(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	if err := c.orderUseCase.CancelOrder(uint(id)); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "注文をキャンセルしました"})
}