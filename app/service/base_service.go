package service

import "gorm.io/gorm"

type BaseService struct {
	UserService *UserService
	// 他のサービスを追加する場合は、ここにフィールドを追加します
	// Example:
	// ProductService *ProductService
}

func NewBaseService(db *gorm.DB) *BaseService {
	return &BaseService{
		UserService: &UserService{DB: db},
		// 他のサービスの初期化を追加します
		// Example:
		// ProductService: &ProductService{DB: db},
	}
}
