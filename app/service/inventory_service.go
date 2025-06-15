package service

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

type InventoryService struct {
	DB *gorm.DB
}

func NewInventoryService(db *gorm.DB) InventoryService {
	return InventoryService{
		DB: db,
	}
}

func (s *InventoryService) GetInventories() ([]model.Inventory, error) {
	var inventories []model.Inventory
	if err := s.DB.Preload("Product").Find(&inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}

func (s *InventoryService) GetInventoryDetail(id uint) (*model.Inventory, error) {
	var inventory model.Inventory
	if err := s.DB.Preload("Product").First(&inventory, id).Error; err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (s *InventoryService) CreateInventory(inventory model.Inventory) error {
	return s.DB.Create(&inventory).Error
}

func (s *InventoryService) UpdateInventory(inventory model.Inventory) error {
	return s.DB.Save(&inventory).Error
}

func (s *InventoryService) DeleteInventory(id uint) error {
	return s.DB.Delete(&model.Inventory{}, id).Error
}

func (s *InventoryService) GetInventoryByProductAndTenant(productID, tenantID uint) (*model.Inventory, error) {
	var inventory model.Inventory
	if err := s.DB.Where("product_id = ? AND tenant_id = ?", productID, tenantID).First(&inventory).Error; err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (s *InventoryService) SaveInventory(inventory *model.Inventory) error {
	return s.DB.Save(inventory).Error
}
