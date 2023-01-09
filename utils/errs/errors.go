package errs

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
)

const (
	PropErr           = "MISSING_REQUEST_PROPERTY"
	PropMess          = "Please set the correct value."
	NotFond           = "MISSING_REQUEST_USERID"
	NotFondMess       = "Non-existent user ID."
	Unauthorized      = "MISSING_REQUEST_UNAUTHORIZED"
	UnauthorizedMesss = "Unauthorized error."
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

func NewUnauthorizedError(t, msg string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Type:    t,
		Message: msg,
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

func ReturnJsonValidation(s interface{}) *AppError {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		str := err.Error()
		i := strings.Index(str, "Error")
		fmt.Println(str[i:])
		return NewBadRequestError(PropErr, str[i:])
	}
	return nil
}
