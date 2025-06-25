package controller

import (
	"myapp/app/usecase"
)

type BaseController struct {
	UserController         UserController
	InventoryController    InventoryController
	ProductController      ProductController
	TenantController       TenantController
	OrderController        OrderController
	PriceSettingController PriceSettingController
}

func NewBaseController(baseUseCase usecase.BaseUseCase) BaseController {
	return BaseController{
		UserController:         NewUserController(baseUseCase.UserUseCase),
		InventoryController:    NewInventoryController(baseUseCase.InventoryUseCase),
		ProductController:      NewProductController(baseUseCase.ProductUseCase),
		TenantController:       NewTenantController(baseUseCase.TenantUseCase),
		OrderController:        NewOrderController(baseUseCase.OrderUseCase),
		PriceSettingController: NewPriceSettingController(baseUseCase.PriceSettingUseCase),
	}
}
