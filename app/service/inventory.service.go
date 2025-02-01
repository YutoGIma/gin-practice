package service

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

type InventoryService struct {
	DB *gorm.DB
}

func (s *InventoryService) GetInventories() ([]model.Inventory, error) {
	var inventories []model.Inventory
	if err := s.DB.Preload("Product").Find(&inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}
