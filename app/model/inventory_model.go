package model

type Inventory struct {
	BaseModel
	ProductID   uint    `json:"product_id" validate:"required"`
	Product     Product `json:"product" gorm:"foreignKey:ProductID"`
	TenantID    uint    `json:"tenant_id" validate:"required"`
	Tenant      Tenant  `json:"tenant" gorm:"foreignKey:TenantID"`
	Quantity    int     `json:"quantity" validate:"required,gte=0"`
	MinQuantity int     `json:"min_quantity"`
	MaxQuantity int     `json:"max_quantity"`
}
