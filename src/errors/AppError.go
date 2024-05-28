package errors

import (
	"fmt"
	"mealwhile/logic/model"
	"net/http"
)

/*

Custom error type for the application

Codes:
400: Bad request
401: Unauthorized
403: Forbidden
404: Not found

*/

type AppError struct {
	Cause   error
	Code    int
	Message string
}

func NewAppError(cause error, code int, message string, translationKey string) AppError {
	return AppError{Cause: cause, Code: code, Message: message}
}

func NewEntityNotFound(entity model.CrudEntity, identifier string) AppError {
	return AppError{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("The %s with %s does not exist", entity.EntityName(), identifier),
	}
}

func NewEntityAlreadyExists(entity model.CrudEntity, identifier string) AppError {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf("The %s with %s already exists", entity.EntityName(), identifier),
	}
}

func NewServerError(message string) error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewBadRequest(message string) error {
	return AppError{
		Code:    http.StatusBadGateway,
		Message: message,
	}
}

func (e AppError) Error() string {
	return e.Message
}
