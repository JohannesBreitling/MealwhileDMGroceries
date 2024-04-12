package data

import (
	"fmt"
	"mealwhile/logic/model"
)

/*
Explanation of the error codes:
- 0: Internal error with the repository
- 1: Entity not found
- 2: Entity already exists
*/
const min_error_code = 0
const max_error_code = 2

type RepositoryError struct {
	Code    int
	Message string
}

func NewRepositoryError(message string, code int) RepositoryError {
	if code > max_error_code || code < min_error_code {
		return RepositoryError{0, "Internal repository error"}
	}

	return RepositoryError{code, message}
}

func (err RepositoryError) Error() string {
	switch err.Code {
	case 0:
		return fmt.Sprintf("Internal repository error: %s", err.Message)
	case 1:
		return fmt.Sprintf("Entity not found: %s", err.Message)
	case 2:
		return fmt.Sprintf("Entity already exists: %s", err.Message)

	default:
		return "Internal repository error"
	}
}

func NewNotFoundRepositoryError(entity model.CrudEntity, attribute string) RepositoryError {
	message := fmt.Sprintf("%s with %s was not found", entity.EntityName(), attribute)
	return NewRepositoryError(message, 1)
}

func NewAlreadyExistsRepositoryError(entity model.CrudEntity, attribute string) RepositoryError {
	message := fmt.Sprintf("%s with %s already exists", entity.EntityName(), attribute)
	return NewRepositoryError(message, 2)
}

func NewDBRepositoryError(message string) RepositoryError {
	return NewRepositoryError(message, 0)
}
