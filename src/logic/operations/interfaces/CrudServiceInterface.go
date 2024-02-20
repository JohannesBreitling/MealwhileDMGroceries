package interfaces

import (
	"mealwhile/logic/model"
)

type CrudServiceInterface interface {
	Create(entity model.CrudEntity) (model.CrudEntity, error)
	Read(entity model.CrudEntity) (model.CrudEntity, error)
	Update(entity model.CrudEntity) (model.CrudEntity, error)
	Delete(entity model.CrudEntity) error
}
