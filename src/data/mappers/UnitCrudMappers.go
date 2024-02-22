package mappers

import (
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"
)

type UnitMapper struct {
}

func (UnitMapper) EntityToPersistenceEntity(e model.CrudEntity) persistenceentites.CrudPersistenceEntity {
	unit := e.(*model.Unit)

	return persistenceentites.UnitPersistenceEntity{
		Id:           unit.Id,
		Name:         unit.Name,
		Abbreviation: unit.Abbreviation,
	}
}

func (UnitMapper) PersistenceEntityToEntity(pe persistenceentites.CrudPersistenceEntity) model.CrudEntity {
	upe := pe.(persistenceentites.UnitPersistenceEntity)

	return &model.Unit{
		Id:           upe.Id,
		Name:         upe.Name,
		Abbreviation: upe.Abbreviation,
	}
}

func (u UnitMapper) PersisteceEntitesToEntites(pes []persistenceentites.CrudPersistenceEntity) []model.CrudEntity {
	var es []model.CrudEntity

	for _, pe := range pes {
		es = append(es, u.PersistenceEntityToEntity(pe))
	}

	return es
}

func (u UnitMapper) EntitesToPersisteceEntites(es []model.CrudEntity) []persistenceentites.CrudPersistenceEntity {
	var pes []persistenceentites.CrudPersistenceEntity

	for _, e := range es {
		pes = append(pes, u.EntityToPersistenceEntity(e))
	}

	return pes
}
