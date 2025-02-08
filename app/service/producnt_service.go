package service

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func (s ProductService) GetProducts() ([]model.Product, error) {
	var products []model.Product
	err := s.DB.Find(products).Error
	return products, err
}

func (s ProductService) GetProductDetail(id int) (model.Product, error) {
	var product model.Product
	err := s.DB.First(product, id).Error
	return product, err
}

func (s ProductService) CreateProduct(product model.Product) error {
	return s.DB.Create(product).Error
}

func (s ProductService) UpdateProduct(product model.Product) error {
	return s.DB.Save(product).Error
}

func (s ProductService) DeleteProduct(id int) error {
	return s.DB.Delete(model.Product{}, id).Error
}
