package mappers

import (
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"
)

type CrudMappersInterface interface {
	EntityToPersistenceEntity(e model.CrudEntity) persistenceentites.CrudPersistenceEntity
	PersistenceEntityToEntity(pe persistenceentites.CrudPersistenceEntity) model.CrudEntity
	PersisteceEntitesToEntites(pes []persistenceentites.CrudPersistenceEntity) []model.CrudEntity
	EntitesToPersisteceEntites(es []model.CrudEntity) []persistenceentites.CrudPersistenceEntity
}
