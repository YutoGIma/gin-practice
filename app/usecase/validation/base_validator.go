package validation

import (
	"myapp/app/errors"
)

// BaseValidator provides common validation methods
type BaseValidator struct{}

// NewBaseValidator creates a new base validator
func NewBaseValidator() BaseValidator {
	return BaseValidator{}
}

// ValidateID validates that an ID is greater than 0
func (v BaseValidator) ValidateID(id uint, fieldName string) error {
	if id == 0 {
		return errors.NewValidationError(fieldName + "は必須です")
	}
	return nil
}

// ValidateRequiredString validates that a string is not empty
func (v BaseValidator) ValidateRequiredString(value, fieldName string) error {
	if value == "" {
		return errors.NewValidationError(fieldName + "は必須です")
	}
	return nil
}

// ValidatePositiveNumber validates that a number is positive
func (v BaseValidator) ValidatePositiveNumber(value int, fieldName string) error {
	if value <= 0 {
		return errors.NewValidationError(fieldName + "は1以上である必要があります")
	}
	return nil
}

// ValidateNonNegativeNumber validates that a number is non-negative
func (v BaseValidator) ValidateNonNegativeNumber(value int, fieldName string) error {
	if value < 0 {
		return errors.NewValidationError(fieldName + "は0以上である必要があります")
	}
	return nil
}
