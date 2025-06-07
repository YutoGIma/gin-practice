package controller

import (
	"myapp/app/errors"
	"myapp/app/model"
	"myapp/app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase *usecase.ProductUseCase
}

func NewProductController(productUseCase *usecase.ProductUseCase) *ProductController {
	return &ProductController{
		productUseCase: productUseCase,
	}
}

func (c *ProductController) Create(ctx *gin.Context) {
	var input model.Product
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.Error(errors.NewValidationError("Invalid input"))
		return
	}

	result, err := c.productUseCase.Create(input)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(201, result)
}

func (c *ProductController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(errors.NewValidationError("Invalid ID"))
		return
	}

	var input model.Product
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.Error(errors.NewValidationError("Invalid input"))
		return
	}

	result, err := c.productUseCase.Update(uint(id), input)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, result)
}

func (c *ProductController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(errors.NewValidationError("Invalid ID"))
		return
	}

	if err := c.productUseCase.Delete(uint(id)); err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(204)
}

func (c *ProductController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(errors.NewValidationError("Invalid ID"))
		return
	}

	result, err := c.productUseCase.GetByID(uint(id))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, result)
}

func (c *ProductController) List(ctx *gin.Context) {
	result, err := c.productUseCase.List()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, result)
}
