package usecase

import (
	"myapp/app/service"
)

type BaseUseCase struct {
	ProductUseCase      ProductUseCase
	TenantUseCase       TenantUseCase
	InventoryUseCase    InventoryUseCase
	UserUseCase         UserUseCase
	OrderUseCase        OrderUseCase
	PriceSettingUseCase PriceSettingUseCase
}

func NewBaseUseCase(baseService service.BaseService) BaseUseCase {
	return BaseUseCase{
		ProductUseCase:      NewProductUseCase(baseService.ProductService),
		TenantUseCase:       NewTenantUseCase(baseService.TenantService),
		InventoryUseCase:    NewInventoryUseCase(baseService.InventoryService),
		UserUseCase:         NewUserUseCase(baseService.UserService),
		OrderUseCase:        NewOrderUseCase(baseService.OrderService, baseService.InventoryService, baseService.ProductService),
		PriceSettingUseCase: NewPriceSettingUseCase(baseService.PriceSettingService, baseService.InventoryService),
	}
}
