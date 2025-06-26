package validation

import (
	"errors"
	"myapp/app/model"
	"regexp"
	"strconv"
)

// ProductValidator handles product-related validations
type ProductValidator struct {
	BaseValidator
}

// NewProductValidator creates a new product validator
func NewProductValidator() ProductValidator {
	return ProductValidator{
		BaseValidator: NewBaseValidator(),
	}
}

// ValidateCreateProduct validates product creation
func (v ProductValidator) ValidateCreateProduct(p model.Product) error {
	if err := v.ValidateRequiredString(p.Name, "商品名"); err != nil {
		return err
	}

	if err := v.ValidateBarcode(p.Barcode); err != nil {
		return err
	}

	if p.CategoryId <= 0 {
		return errors.New("カテゴリIDは必須です")
	}

	if p.PurchasePrice < 0 {
		return errors.New("仕入れ価格は0以上である必要があります")
	}

	return nil
}

// ValidateBarcode validates JAN-13 barcode
func (v ProductValidator) ValidateBarcode(barcode string) error {
	// 13桁の数字かチェック
	match, _ := regexp.MatchString(`^[0-9]{13}$`, barcode)
	if !match {
		return errors.New("バーコードは13桁の数字で入力してください")
	}

	// チェックディジットの検証
	sum := 0
	for i := 0; i < 12; i++ {
		digit, _ := strconv.Atoi(string(barcode[i]))
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}
	checkDigit := (10 - (sum % 10)) % 10
	lastDigit, _ := strconv.Atoi(string(barcode[12]))
	if checkDigit != lastDigit {
		return errors.New("正しいバーコードを入力してください")
	}

	return nil
}
