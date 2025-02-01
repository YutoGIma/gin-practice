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
