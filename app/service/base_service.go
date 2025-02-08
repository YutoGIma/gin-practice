package service

import "gorm.io/gorm"

type BaseService struct {
	UserService      UserService
	InventoryService InventoryService
	ProductService   ProductService
	// 他のサービスを追加する場合は、ここにフィールドを追加します
	// Example:
	// ProductService *ProductService
}

func NewBaseService(db *gorm.DB) BaseService {
	return BaseService{
		UserService:      UserService{DB: db},
		InventoryService: InventoryService{DB: db},
		ProductService:   ProductService{DB: db},
		// 他のサービスの初期化を追加します
		// Example:
		// ProductService: &ProductService{DB: db},
	}
}
