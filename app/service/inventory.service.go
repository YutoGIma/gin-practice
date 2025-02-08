package service

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

type InventoryService struct {
	DB *gorm.DB
}

func (s InventoryService) GetInventories() ([]model.Inventory, error) {
	var inventories []model.Inventory
	if err := s.DB.Preload("Product").Find(inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}

func (s InventoryService) CreateInventory(inventory model.Inventory) error {
	return s.DB.Create(inventory).Error
}

func (s InventoryService) UpdateInventory(inventory model.Inventory) error {
	return s.DB.Save(inventory).Error
}

func (s InventoryService) DeleteInventory(id uint) error {
	return s.DB.Delete(model.Inventory{}, id).Error
}
