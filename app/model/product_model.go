package model

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	BaseModel
	Name          string  `json:"name" validate:"required"`
	Barcode       string  `json:"barcode" validate:"required,jan13" gorm:"uniqueIndex"`
	CategoryId    int     `json:"master_item_type_id" validate:"required"`
	PurchasePrice float64 `json:"purchase_price" validate:"required,gte=0"`
}

// ValidateJAN13 JANコード(13桁)のバリデーション
func ValidateJAN13(fl validator.FieldLevel) bool {
	barcode := fl.Field().String()
	// 13桁の数字かチェック
	match, _ := regexp.MatchString(`^[0-9]{13}$`, barcode)
	if !match {
		return false
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

	return checkDigit == lastDigit
}
