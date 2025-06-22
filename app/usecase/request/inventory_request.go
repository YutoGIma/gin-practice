package request

import "errors"

type InventoryPurchaseRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	TenantID  uint `json:"tenant_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

func (r InventoryPurchaseRequest) Validate() error {
	if r.Quantity <= 0 {
		return errors.New("数量は1以上である必要があります")
	}
	return nil
}

type InventoryRestockRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	TenantID  uint `json:"tenant_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
	Note      string `json:"note"`
}

type InventoryUpdateRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	TenantID  uint `json:"tenant_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=0"`
}