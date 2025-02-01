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
	err := s.DB.Find(&inventories).Error
	return inventories, err
}
