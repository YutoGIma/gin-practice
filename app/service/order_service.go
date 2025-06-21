package service

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) OrderService {
	return OrderService{db: db}
}

func (s *OrderService) GetOrders() ([]model.Order, error) {
	var orders []model.Order
	err := s.db.Preload("User").Preload("Tenant").Preload("OrderItems.Product").Find(&orders).Error
	return orders, err
}

func (s *OrderService) GetOrderByID(id uint) (*model.Order, error) {
	var order model.Order
	err := s.db.Preload("User").Preload("Tenant").Preload("OrderItems.Product").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *OrderService) GetOrdersByUserID(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := s.db.Preload("User").Preload("Tenant").Preload("OrderItems.Product").
		Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	return s.db.Create(order).Error
}

func (s *OrderService) UpdateOrderStatus(id uint, status model.OrderStatus) error {
	return s.db.Model(&model.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (s *OrderService) BeginTransaction() *gorm.DB {
	return s.db.Begin()
}

func (s *OrderService) CreateOrderWithTx(tx *gorm.DB, order *model.Order) error {
	return tx.Create(order).Error
}