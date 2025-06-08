package usecase

import (
	"myapp/app/service"
)

type BaseUseCase struct {
	ProductUseCase   ProductUseCase
	TenantUseCase    TenantUseCase
	InventoryUseCase InventoryUseCase
	UserUseCase      UserUseCase
}

func NewBaseUseCase(baseService service.BaseService) BaseUseCase {
	return BaseUseCase{
		ProductUseCase:   NewProductUseCase(baseService.ProductService),
		TenantUseCase:    NewTenantUseCase(baseService.TenantService),
		InventoryUseCase: NewInventoryUseCase(baseService.InventoryService),
		UserUseCase:      NewUserUseCase(baseService.UserService),
	}
}
