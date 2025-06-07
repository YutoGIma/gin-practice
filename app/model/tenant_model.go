package model

import (
	"time"
)

// Tenant は拠点（テナント）を表すモデル
type Tenant struct {
	BaseModel
	Name        string     `json:"name" validate:"required"`
	Code        string     `json:"code" validate:"required" gorm:"uniqueIndex"`
	Address     string     `json:"address" validate:"required"`
	PhoneNumber string     `json:"phone_number" validate:"required"`
	Email       string     `json:"email" validate:"required,email"`
	IsActive    bool       `json:"is_active" gorm:"default:true"`
	OpenedAt    time.Time  `json:"opened_at"`
	ClosedAt    *time.Time `json:"closed_at,omitempty"`
}

// GetModels にテナントモデルを追加
func init() {
	models = append(models, &Tenant{})
}
