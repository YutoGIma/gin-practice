package request

import "time"

type CreatePriceSettingRequest struct {
	InventoryID uint       `json:"inventory_id" binding:"required"`
	Price       float64    `json:"price" binding:"required,gt=0"`
	SalePrice   *float64   `json:"sale_price,omitempty"`
	StartDate   time.Time  `json:"start_date" binding:"required"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	Note        string     `json:"note,omitempty"`
}

type UpdatePriceSettingRequest struct {
	Price     *float64    `json:"price,omitempty"`
	SalePrice **float64   `json:"sale_price,omitempty"`
	StartDate *time.Time  `json:"start_date,omitempty"`
	EndDate   **time.Time `json:"end_date,omitempty"`
	IsActive  *bool       `json:"is_active,omitempty"`
	Note      *string     `json:"note,omitempty"`
}
