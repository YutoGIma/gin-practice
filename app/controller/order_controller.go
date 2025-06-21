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

func (c *OrderController) GetOrders(ctx *gin.Context) {
	orders, err := c.orderUseCase.GetOrders()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

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

func (c *OrderController) GetUserOrders(ctx *gin.Context) {
	idStr := ctx.Param("user_id")
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