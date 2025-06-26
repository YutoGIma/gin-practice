package model

import (
	"time"
)

type PriceSetting struct {
	BaseModel
	InventoryID uint       `gorm:"not null;index:idx_inventory_date"`
	Inventory   Inventory  `gorm:"foreignKey:InventoryID"`
	Price       float64    `gorm:"not null"`
	SalePrice   *float64   `gorm:"default:null"`
	StartDate   time.Time  `gorm:"not null;index:idx_inventory_date"`
	EndDate     *time.Time `gorm:"index:idx_inventory_date"`
	IsActive    bool       `gorm:"not null;default:true"`
	Note        string     `gorm:"type:text"`
}

func (PriceSetting) TableName() string {
	return "price_settings"
}

// IsCurrentlyActive checks if the price setting is active for the current time
func (p *PriceSetting) IsCurrentlyActive() bool {
	if !p.IsActive {
		return false
	}

	now := time.Now()

	// Check if current time is after start date
	if now.Before(p.StartDate) {
		return false
	}

	// Check if end date is set and current time is after end date
	if p.EndDate != nil && now.After(*p.EndDate) {
		return false
	}

	return true
}

// GetEffectivePrice returns the current effective price (sale price if available, otherwise regular price)
func (p *PriceSetting) GetEffectivePrice() float64 {
	if p.SalePrice != nil && p.IsCurrentlyActive() {
		return *p.SalePrice
	}
	return p.Price
}

// IsOnSale checks if the item is currently on sale
func (p *PriceSetting) IsOnSale() bool {
	return p.SalePrice != nil && p.IsCurrentlyActive()
}
