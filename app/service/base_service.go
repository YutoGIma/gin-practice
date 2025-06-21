package service

import (
	"gorm.io/gorm"
)

type BaseService struct {
	ProductService   ProductService
	TenantService    TenantService
	InventoryService InventoryService
	UserService      UserService
	OrderService     OrderService
}

func NewBaseService(db *gorm.DB) BaseService {
	return BaseService{
		ProductService:   NewProductService(db),
		TenantService:    NewTenantService(db),
		InventoryService: NewInventoryService(db),
		UserService:      NewUserService(db),
		OrderService:     NewOrderService(db),
	}
}
