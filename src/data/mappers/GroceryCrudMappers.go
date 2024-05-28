package mappers

import (
	"encoding/json"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"
)

type GroceryMapper struct {
}

func (GroceryMapper) EntityToPersistenceEntity(e model.CrudEntity) persistenceentites.CrudPersistenceEntity {
	grocery := e.(*model.Grocery)

	flagIds, _ := json.Marshal(grocery.FlagIds)

	return persistenceentites.GroceryPersistenceEntity{
		Id:      grocery.Id,
		Name:    grocery.Name,
		FlagIds: string(flagIds),
	}
}

func (GroceryMapper) PersistenceEntityToEntity(pe persistenceentites.CrudPersistenceEntity) model.CrudEntity {
	gpe := pe.(persistenceentites.GroceryPersistenceEntity)

	var flags []string
	json.Unmarshal([]byte(gpe.FlagIds), &flags)

	return &model.Grocery{
		Id:      gpe.Id,
		Name:    gpe.Name,
		FlagIds: flags,
	}
}

func (u GroceryMapper) PersisteceEntitesToEntites(pes []persistenceentites.CrudPersistenceEntity) []model.CrudEntity {
	var es []model.CrudEntity

	for _, pe := range pes {
		es = append(es, u.PersistenceEntityToEntity(pe))
	}

	return es
}

func (u GroceryMapper) EntitesToPersisteceEntites(es []model.CrudEntity) []persistenceentites.CrudPersistenceEntity {
	var pes []persistenceentites.CrudPersistenceEntity

	for _, e := range es {
		pes = append(pes, u.EntityToPersistenceEntity(e))
	}

	return pes
}
