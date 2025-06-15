package request

import (
	"errors"
)

type InventoryPurchaseRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	TenantID  uint `json:"tenant_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,gt=0"`
}

func (r *InventoryPurchaseRequest) Validate() error {
	if r.ProductID == 0 {
		return errors.New("product_id is required")
	}
	if r.TenantID == 0 {
		return errors.New("tenant_id is required")
	}
	if r.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	return nil
}
