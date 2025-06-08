package usecase

import (
	"myapp/app/errors"
	"myapp/app/model"
	"myapp/app/service"
)

type InventoryUseCase struct {
	inventoryService service.InventoryService
}

func NewInventoryUseCase(inventoryService service.InventoryService) InventoryUseCase {
	return InventoryUseCase{
		inventoryService: inventoryService,
	}
}

func (uc *InventoryUseCase) List() ([]model.Inventory, error) {
	inventories, err := uc.inventoryService.GetInventories()
	if err != nil {
		return nil, errors.NewInternalError("Failed to list inventories", err)
	}
	return inventories, nil
}

func (uc *InventoryUseCase) GetByID(id uint) (*model.Inventory, error) {
	inventory, err := uc.inventoryService.GetInventoryDetail(id)
	if err != nil {
		return nil, errors.NewInternalError("Failed to get inventory", err)
	}
	return inventory, nil
}

func (uc *InventoryUseCase) Create(input model.Inventory) (*model.Inventory, error) {
	if err := uc.inventoryService.CreateInventory(input); err != nil {
		return nil, errors.NewInternalError("Failed to create inventory", err)
	}
	return &input, nil
}

func (uc *InventoryUseCase) Update(id uint, input model.Inventory) (*model.Inventory, error) {
	input.ID = id
	if err := uc.inventoryService.UpdateInventory(input); err != nil {
		return nil, errors.NewInternalError("Failed to update inventory", err)
	}
	return &input, nil
}

func (uc *InventoryUseCase) Delete(id uint) error {
	if err := uc.inventoryService.DeleteInventory(id); err != nil {
		return errors.NewInternalError("Failed to delete inventory", err)
	}
	return nil
}
