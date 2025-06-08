package controller

import (
	"myapp/app/model"
	"myapp/app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TenantController struct {
	tenantUseCase usecase.TenantUseCase
}

func NewTenantController(tenantUseCase usecase.TenantUseCase) TenantController {
	return TenantController{
		tenantUseCase: tenantUseCase,
	}
}

func (c TenantController) GetTenants(ctx *gin.Context) {
	tenants, err := c.tenantUseCase.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tenants)
}

// func (c TenantController) GetTenantDetail(ctx *gin.Context) {
// 	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	tenant, err := c.tenantUseCase.GetTenantDetail(uint(id))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, tenant)
// }

func (c TenantController) CreateTenant(ctx *gin.Context) {
	var tenant model.Tenant
	if err := ctx.ShouldBindJSON(&tenant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTenant, err := c.tenantUseCase.Create(tenant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, createdTenant)
}

func (c TenantController) UpdateTenant(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var tenant model.Tenant
	if err := ctx.ShouldBindJSON(&tenant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTenant, err := c.tenantUseCase.Update(uint(id), tenant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedTenant)
}

func (c TenantController) DeleteTenant(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.tenantUseCase.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Tenant deleted successfully"})
}
