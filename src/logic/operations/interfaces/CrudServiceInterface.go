package interfaces

import (
	"mealwhile/logic/model"
)

type CrudServiceInterface interface {
	Create(entity model.CrudEntity) (model.CrudEntity, error)
	ReadAll(entity model.CrudEntity) ([]model.CrudEntity, error)
	// Read(id string) (model.CrudEntity, error)
	// Update(entity model.CrudEntity) (model.CrudEntity, error)
	// Delete(id string) error
	// Exists(id string) (bool, error)
}
