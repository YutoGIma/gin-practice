package usecase

import (
	"gorm.io/gorm"
	"myapp/app/errors"
	"myapp/app/model"
)

type ProductUseCase struct {
	db *gorm.DB
}

func NewProductUseCase(db *gorm.DB) *ProductUseCase {
	return &ProductUseCase{
		db: db,
	}
}

func (uc *ProductUseCase) Create(input model.Product) (*model.Product, error) {
	if err := model.ValidateCreateProduct(input); err != nil {
		return nil, errors.NewValidationError(err.Error())
	}

	if err := uc.db.Create(&input).Error; err != nil {
		return nil, errors.NewInternalError("Failed to create product", err)
	}

	return &input, nil
}

func (uc *ProductUseCase) Update(id uint, input model.Product) (*model.Product, error) {
	var product model.Product
	if err := uc.db.First(&product, id).Error; err != nil {
		return nil, errors.NewNotFoundError("Product not found")
	}

	if err := model.ValidateCreateProduct(input); err != nil {
		return nil, errors.NewValidationError(err.Error())
	}

	if err := uc.db.Model(&product).Updates(input).Error; err != nil {
		return nil, errors.NewInternalError("Failed to update product", err)
	}

	return &product, nil
}

func (uc *ProductUseCase) Delete(id uint) error {
	if err := uc.db.Delete(&model.Product{}, id).Error; err != nil {
		return errors.NewInternalError("Failed to delete product", err)
	}
	return nil
}

func (uc *ProductUseCase) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := uc.db.First(&product, id).Error; err != nil {
		return nil, errors.NewNotFoundError("Product not found")
	}
	return &product, nil
}

func (uc *ProductUseCase) List() ([]model.Product, error) {
	var products []model.Product
	if err := uc.db.Find(&products).Error; err != nil {
		return nil, errors.NewInternalError("Failed to list products", err)
	}
	return products, nil
}
