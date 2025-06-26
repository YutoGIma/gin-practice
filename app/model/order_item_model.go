package model

type OrderItem struct {
	BaseModel
	OrderID   uint    `gorm:"not null"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Subtotal  float64 `gorm:"not null"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
