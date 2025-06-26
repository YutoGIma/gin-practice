package validation

import (
	"fmt"
	"myapp/app/errors"
	"myapp/app/usecase/request"
)

// InventoryValidator handles inventory-related validations
type InventoryValidator struct {
	BaseValidator
}

// NewInventoryValidator creates a new inventory validator
func NewInventoryValidator() InventoryValidator {
	return InventoryValidator{
		BaseValidator: NewBaseValidator(),
	}
}

// ValidatePurchaseRequest validates inventory purchase request
func (v InventoryValidator) ValidatePurchaseRequest(req request.InventoryPurchaseRequest) error {
	if err := v.ValidateID(req.ProductID, "商品ID"); err != nil {
		return err
	}

	if err := v.ValidateID(req.TenantID, "テナントID"); err != nil {
		return err
	}

	if err := v.ValidatePositiveNumber(req.Quantity, "購入数量"); err != nil {
		return err
	}

	return nil
}

// ValidateRestockRequest validates inventory restock request
func (v InventoryValidator) ValidateRestockRequest(req request.InventoryRestockRequest) error {
	if err := v.ValidateID(req.ProductID, "商品ID"); err != nil {
		return err
	}

	if err := v.ValidateID(req.TenantID, "テナントID"); err != nil {
		return err
	}

	if err := v.ValidatePositiveNumber(req.Quantity, "入荷数量"); err != nil {
		return err
	}

	return nil
}

// ValidateUpdateRequest validates inventory update request
func (v InventoryValidator) ValidateUpdateRequest(req request.InventoryUpdateRequest) error {
	if err := v.ValidateID(req.ProductID, "商品ID"); err != nil {
		return err
	}

	if err := v.ValidateID(req.TenantID, "テナントID"); err != nil {
		return err
	}

	if err := v.ValidateNonNegativeNumber(req.Quantity, "在庫数量"); err != nil {
		return err
	}

	return nil
}

// ValidateSufficientInventory validates that inventory has sufficient quantity
func (v InventoryValidator) ValidateSufficientInventory(currentQuantity, requestedQuantity int, productName string) error {
	if currentQuantity < requestedQuantity {
		return errors.NewValidationError(fmt.Sprintf("商品「%s」の在庫が不足しています。現在の在庫: %d", productName, currentQuantity))
	}
	return nil
}
