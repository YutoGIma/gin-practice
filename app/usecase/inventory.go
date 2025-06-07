package usecase

import (
	"gorm.io/gorm"
	"myapp/app/errors"
	"myapp/app/model"
)

type InventoryUseCase struct {
	db *gorm.DB
}

func NewInventoryUseCase(db *gorm.DB) *InventoryUseCase {
	return &InventoryUseCase{
		db: db,
	}
}

func (uc *InventoryUseCase) Create(input model.Inventory) (*model.Inventory, error) {
	if err := uc.db.Create(&input).Error; err != nil {
		return nil, errors.NewInternalError("Failed to create inventory", err)
	}
	return &input, nil
}

func (uc *InventoryUseCase) Update(id uint, input model.Inventory) (*model.Inventory, error) {
	var inventory model.Inventory
	if err := uc.db.First(&inventory, id).Error; err != nil {
		return nil, errors.NewNotFoundError("Inventory not found")
	}

	if err := uc.db.Model(&inventory).Updates(input).Error; err != nil {
		return nil, errors.NewInternalError("Failed to update inventory", err)
	}

	return &inventory, nil
}

func (uc *InventoryUseCase) Delete(id uint) error {
	if err := uc.db.Delete(&model.Inventory{}, id).Error; err != nil {
		return errors.NewInternalError("Failed to delete inventory", err)
	}
	return nil
}

func (uc *InventoryUseCase) GetByID(id uint) (*model.Inventory, error) {
	var inventory model.Inventory
	if err := uc.db.First(&inventory, id).Error; err != nil {
		return nil, errors.NewNotFoundError("Inventory not found")
	}
	return &inventory, nil
}

func (uc *InventoryUseCase) List() ([]model.Inventory, error) {
	var inventories []model.Inventory
	if err := uc.db.Find(&inventories).Error; err != nil {
		return nil, errors.NewInternalError("Failed to list inventories", err)
	}
	return inventories, nil
}

func (uc *InventoryUseCase) GetByProductID(productID uint) (*model.Inventory, error) {
	var inventory model.Inventory
	if err := uc.db.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		return nil, errors.NewNotFoundError("Inventory not found for product")
	}
	return &inventory, nil
}

func (uc *InventoryUseCase) GetByTenantID(tenantID uint) ([]model.Inventory, error) {
	var inventories []model.Inventory
	if err := uc.db.Where("tenant_id = ?", tenantID).Find(&inventories).Error; err != nil {
		return nil, errors.NewInternalError("Failed to list inventories for tenant", err)
	}
	return inventories, nil
}
