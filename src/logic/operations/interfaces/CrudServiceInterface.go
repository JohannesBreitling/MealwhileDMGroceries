package interfaces

import (
	"mealwhile/logic/model"
)

type CrudServiceInterface interface {
	Create(entity model.CrudEntity) (model.CrudEntity, error)
	ReadAll(target model.CrudEntity) ([]model.CrudEntity, error)
	Read(target model.CrudEntity, id string) (model.CrudEntity, error)
	Update(entity model.CrudEntity) (model.CrudEntity, error)
	Delete(target model.CrudEntity, id string) error
}
