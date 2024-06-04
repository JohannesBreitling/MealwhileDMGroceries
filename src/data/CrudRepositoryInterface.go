package data

import (
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"
)

type CrudRepositoryInterface interface {
	Create(entity model.CrudEntity) (model.CrudEntity, error)
	ReadAll() ([]model.CrudEntity, error)
	Read(id string) (model.CrudEntity, error)
	Update(entity model.CrudEntity) (model.CrudEntity, error)
	Delete(id string) error
	Exists(id string) (bool, error)
	ReadPe(id string) (*persistenceentites.CrudPersistenceEntity, error)
}
