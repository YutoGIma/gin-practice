package model

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	BaseModel
	UserID      uint        `gorm:"not null"`
	User        User        `gorm:"foreignKey:UserID"`
	TenantID    uint        `gorm:"not null"`
	Tenant      Tenant      `gorm:"foreignKey:TenantID"`
	TotalAmount float64     `gorm:"not null"`
	Status      OrderStatus `gorm:"type:varchar(20);not null;default:'pending'"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID"`
	OrderedAt   time.Time   `gorm:"not null"`
}

func (Order) TableName() string {
	return "orders"
}