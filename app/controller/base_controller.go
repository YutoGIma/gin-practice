package controller

import (
	"myapp/app/usecase"
)

type BaseController struct {
	UserController      UserController
	InventoryController InventoryController
	ProductController   ProductController
	TenantController    TenantController
}

func NewBaseController(baseUseCase usecase.BaseUseCase) BaseController {
	return BaseController{
		UserController:      NewUserController(baseUseCase.UserUseCase),
		InventoryController: NewInventoryController(baseUseCase.InventoryUseCase),
		ProductController:   NewProductController(baseUseCase.ProductUseCase),
		TenantController:    NewTenantController(baseUseCase.TenantUseCase),
	}
}
