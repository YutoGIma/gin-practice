package model

type Inventory struct {
	BaseModel
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
	// StoreId     int `json:"store_id"`
	MinQuantity int `json:"min_quantity"`
	MaxQuantity int `json:"max_quantity"`

	Product Product `json:"product" gorm:"foreignKey:ProductId"`
}
