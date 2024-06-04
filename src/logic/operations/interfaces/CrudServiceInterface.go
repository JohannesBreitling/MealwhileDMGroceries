package interfaces

import (
	"mealwhile/logic/model"
	"mealwhile/logic/model/requests"
)

type CrudServiceInterface interface {
	Create(request requests.CrudRequest) (model.CrudEntity, error)
	ReadAll() ([]model.CrudEntity, error)
	Read(id string) (model.CrudEntity, error)
	Update(request requests.CrudRequest) (model.CrudEntity, error)
	Delete(id string) error
}
