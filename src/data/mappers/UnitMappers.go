package mappers

import (
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"
)

func UnitToUnitPersistenceEntity(unit model.Unit) persistenceentites.UnitPersistenceEntity {
	return persistenceentites.UnitPersistenceEntity{
		Id:           unit.Id,
		Name:         unit.Name,
		Abbreviation: unit.Abbreviation,
	}
}
