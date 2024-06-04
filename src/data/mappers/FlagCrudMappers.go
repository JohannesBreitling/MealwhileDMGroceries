package mappers

import (
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"
	"reflect"
)

type FlagMapper struct {
}

func (FlagMapper) EntityToPersistenceEntity(e model.CrudEntity) persistenceentites.CrudPersistenceEntity {
	flag := e.(*model.Flag)

	return persistenceentites.FlagPersistenceEntity{
		Id:          flag.Id,
		Name:        flag.Name,
		Description: flag.Description,
	}
}

func (FlagMapper) PersistenceEntityToEntity(pe persistenceentites.CrudPersistenceEntity) model.CrudEntity {
	var fpe persistenceentites.FlagPersistenceEntity
	switch reflect.TypeOf(pe) {
	case reflect.TypeOf(persistenceentites.FlagPersistenceEntity{}):
		fpe = pe.(persistenceentites.FlagPersistenceEntity)
	case reflect.TypeOf(&persistenceentites.FlagPersistenceEntity{}):
		fpe = *(pe.(*persistenceentites.FlagPersistenceEntity))
	default:
		fpe = persistenceentites.FlagPersistenceEntity{}
	}

	return &model.Flag{
		Id:          fpe.Id,
		Name:        fpe.Name,
		Description: fpe.Description,
	}
}

func (u FlagMapper) PersisteceEntitesToEntites(pes []persistenceentites.CrudPersistenceEntity) []model.CrudEntity {
	var es []model.CrudEntity

	for _, pe := range pes {
		es = append(es, u.PersistenceEntityToEntity(pe))
	}

	return es
}

func (u FlagMapper) EntitesToPersisteceEntites(es []model.CrudEntity) []persistenceentites.CrudPersistenceEntity {
	var pes []persistenceentites.CrudPersistenceEntity

	for _, e := range es {
		pes = append(pes, u.EntityToPersistenceEntity(e))
	}

	return pes
}
