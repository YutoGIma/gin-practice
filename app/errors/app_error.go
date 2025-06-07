package errors

import "fmt"

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewValidationError(message string) *AppError {
	return NewAppError(400, message, nil)
}

func NewNotFoundError(message string) *AppError {
	return NewAppError(404, message, nil)
}

func NewInternalError(message string, err error) *AppError {
	return NewAppError(500, message, err)
}
