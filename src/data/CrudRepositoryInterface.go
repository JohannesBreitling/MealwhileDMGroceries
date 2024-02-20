package data

import (
	"mealwhile/logic/model"
)

type CrudRepositoryInterface interface {
	Create(entity model.CrudEntity) (model.CrudEntity, error)
	// Read(entity persistenceentites.CrudPersistenceEntity)
	//Update(entity persistenceentites.CrudPersistenceEntity) (persistenceentites.CrudPersistenceEntity, error)
	//Delete(entity persistenceentites.CrudPersistenceEntity) error
}
