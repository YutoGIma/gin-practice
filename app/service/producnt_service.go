package service

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
	var products []model.Product
	err := s.DB.Find(&products).Error
	return products, err
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	return s.DB.Create(product).Error
}
