package errs

import "net/http"

type AppError struct {
	StatusCode int    `json:",omitempty"`
	Message    string `json:"message"`
}

func (err AppError) AsMessage() *AppError {
	return &AppError{
		Message: err.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}
