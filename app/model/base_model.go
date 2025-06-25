package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel は共通のフィールドを定義します
type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// GetModels マイグレーション用のモデルリストを返します
func GetModels() []interface{} {
	return []interface{}{
		User{},
		Inventory{},
		Product{},
		Tenant{},
		Order{},
		OrderItem{},
		PriceSetting{},
		// 他のモデルを追加する場合は、ここに追加します
	}
}
