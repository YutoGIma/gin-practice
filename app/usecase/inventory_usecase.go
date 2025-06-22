package usecase

import (
	"myapp/app/errors"
	"myapp/app/model"
	"myapp/app/service"
	"myapp/app/usecase/request"
	"myapp/app/usecase/validation"
)

type InventoryUseCase struct {
	inventoryService service.InventoryService
	validator        validation.InventoryValidator
}

func NewInventoryUseCase(inventoryService service.InventoryService) InventoryUseCase {
	return InventoryUseCase{
		inventoryService: inventoryService,
		validator:        validation.NewInventoryValidator(),
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

func (uc *InventoryUseCase) Update(id uint, req request.InventoryUpdateRequest) (*model.Inventory, error) {
	// リクエストのバリデーション
	if err := uc.validator.ValidateUpdateRequest(req); err != nil {
		return nil, err
	}

	// 既存の在庫を取得
	existingInventory, err := uc.inventoryService.GetInventoryDetail(id)
	if err != nil {
		return nil, errors.NewNotFoundError("在庫が見つかりません")
	}

	// 在庫情報を更新
	existingInventory.ProductID = req.ProductID
	existingInventory.TenantID = req.TenantID
	existingInventory.Quantity = req.Quantity

	// 在庫を保存
	if err := uc.inventoryService.SaveInventory(existingInventory); err != nil {
		return nil, errors.NewInternalError("在庫の更新に失敗しました", err)
	}

	// 更新された在庫を再取得（関連データを含む）
	updatedInventory, err := uc.inventoryService.GetInventoryDetail(id)
	if err != nil {
		return nil, errors.NewInternalError("更新した在庫の取得に失敗しました", err)
	}

	return updatedInventory, nil
}

func (uc *InventoryUseCase) Delete(id uint) error {
	if err := uc.inventoryService.DeleteInventory(id); err != nil {
		return errors.NewInternalError("Failed to delete inventory", err)
	}
	return nil
}

func (uc *InventoryUseCase) UpdateOnPurchase(req request.InventoryPurchaseRequest) error {
	// リクエストのバリデーション
	if err := uc.validator.ValidatePurchaseRequest(req); err != nil {
		return err
	}

	// 1. 対象のInventoryを取得
	inventory, err := uc.inventoryService.GetInventoryByProductAndTenant(req.ProductID, req.TenantID)
	if err != nil {
		return errors.NewNotFoundError("Inventory not found")
	}

	// 2. 在庫数のバリデーション
	if inventory.Quantity < req.Quantity {
		return errors.NewValidationError("Insufficient inventory")
	}

	// 3. 在庫数を更新
	inventory.Quantity -= req.Quantity
	if err := uc.inventoryService.SaveInventory(inventory); err != nil {
		return errors.NewInternalError("Failed to update inventory", err)
	}

	return nil
}

func (uc *InventoryUseCase) RestockInventory(req request.InventoryRestockRequest) (*model.Inventory, error) {
	// リクエストのバリデーション
	if err := uc.validator.ValidateRestockRequest(req); err != nil {
		return nil, err
	}
	// 1. 対象のInventoryを取得
	inventory, err := uc.inventoryService.GetInventoryByProductAndTenant(req.ProductID, req.TenantID)
	if err != nil {
		// 在庫レコードが存在しない場合は新規作成
		newInventory := model.Inventory{
			ProductID: req.ProductID,
			TenantID:  req.TenantID,
			Quantity:  req.Quantity,
		}
		if err := uc.inventoryService.CreateInventory(newInventory); err != nil {
			return nil, errors.NewInternalError("在庫の作成に失敗しました", err)
		}
		// 作成した在庫を再取得（関連データを含む）
		createdInventory, err := uc.inventoryService.GetInventoryByProductAndTenant(req.ProductID, req.TenantID)
		if err != nil {
			return nil, errors.NewInternalError("作成した在庫の取得に失敗しました", err)
		}
		return createdInventory, nil
	}

	// 2. 在庫数を増加
	inventory.Quantity += req.Quantity
	if err := uc.inventoryService.SaveInventory(inventory); err != nil {
		return nil, errors.NewInternalError("在庫の更新に失敗しました", err)
	}

	// 3. 更新した在庫を再取得（関連データを含む）
	updatedInventory, err := uc.inventoryService.GetInventoryDetail(inventory.ID)
	if err != nil {
		return nil, errors.NewInternalError("更新した在庫の取得に失敗しました", err)
	}

	return updatedInventory, nil
}
