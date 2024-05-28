package mappers

import (
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"
)

type RecipeMapper struct {
}

func (RecipeMapper) EntityToPersistenceEntity(e model.CrudEntity) persistenceentites.CrudPersistenceEntity {
	flag := e.(*model.Flag)

	return persistenceentites.FlagPersistenceEntity{
		Id:          flag.Id,
		Name:        flag.Name,
		Description: flag.Description,
	}
}

func (RecipeMapper) PersistenceEntityToEntity(pe persistenceentites.CrudPersistenceEntity) model.CrudEntity {
	fpe := pe.(persistenceentites.FlagPersistenceEntity)

	return &model.Flag{
		Id:          fpe.Id,
		Name:        fpe.Name,
		Description: fpe.Description,
	}
}

func (u RecipeMapper) PersisteceEntitesToEntites(pes []persistenceentites.CrudPersistenceEntity) []model.CrudEntity {
	var es []model.CrudEntity

	for _, pe := range pes {
		es = append(es, u.PersistenceEntityToEntity(pe))
	}

	return es
}

func (u RecipeMapper) EntitesToPersisteceEntites(es []model.CrudEntity) []persistenceentites.CrudPersistenceEntity {
	var pes []persistenceentites.CrudPersistenceEntity

	for _, e := range es {
		pes = append(pes, u.EntityToPersistenceEntity(e))
	}

	return pes
}
