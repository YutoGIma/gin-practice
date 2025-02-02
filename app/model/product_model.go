package model

import (
	"errors"
	"regexp"
	"strconv"
)

type Product struct {
	BaseModel
	Name          string  `json:"name" validate:"required"`
	Barcode       string  `json:"barcode" validate:"required,jan13" gorm:"uniqueIndex"`
	CategoryId    int     `json:"master_item_type_id" validate:"required"`
	PurchasePrice float64 `json:"purchase_price" validate:"required,gte=0"`
}

func ValidateCreateProduct(p Product) error {
	// 13桁の数字かチェック
	match, _ := regexp.MatchString(`^[0-9]{13}$`, p.Barcode)
	if !match {
		return errors.New("バーコードは13桁の数字で入力してください")
	}

	// チェックディジットの検証
	sum := 0
	for i := 0; i < 12; i++ {
		digit, _ := strconv.Atoi(string(p.Barcode[i]))
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}
	checkDigit := (10 - (sum % 10)) % 10
	lastDigit, _ := strconv.Atoi(string(p.Barcode[12]))
	if checkDigit != lastDigit {
		return errors.New("正しいバーコードを入力してください")
	}
	return nil
}
