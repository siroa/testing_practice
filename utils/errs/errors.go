package errs

import "net/http"

const (
	PropErr     = "MISSING_REQUEST_PROPERTY"
	PropMess    = "Please set the correct value."
	NotFond     = "MISSING_REQUEST_USERID"
	NotFondMess = "Non-existent user ID."
)

type AppError struct {
	Code    int    `json:",omitempty"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Type:    e.Type,
		Message: e.Message,
	}
}

func NewNotFoundError(t, msg string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Type:    t,
		Message: msg,
	}
}

func NewBadRequestError(t, msg string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Type:    t,
		Message: msg,
	}
}

func NewUnexpectedError(t, msg string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Type:    t,
		Message: msg,
	}
}
