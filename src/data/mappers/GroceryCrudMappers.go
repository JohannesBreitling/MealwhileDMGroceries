package mappers

import (
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"
	"reflect"
)

type GroceryMapper struct {
	flagMapper *FlagMapper
}

func NewGroceryMapper(flagMapper *FlagMapper) GroceryMapper {
	return GroceryMapper{flagMapper: flagMapper}
}

func (gm GroceryMapper) EntityToPersistenceEntity(e model.CrudEntity) persistenceentites.CrudPersistenceEntity {
	grocery := e.(*model.Grocery)

	var flagPes []persistenceentites.FlagPersistenceEntity = []persistenceentites.FlagPersistenceEntity{}
	for _, flag := range grocery.Flags {
		flagPes = append(flagPes, gm.flagMapper.EntityToPersistenceEntity(&flag).(persistenceentites.FlagPersistenceEntity))
	}

	return persistenceentites.GroceryPersistenceEntity{
		Id:    grocery.Id,
		Name:  grocery.Name,
		Flags: flagPes,
	}
}

func (gm GroceryMapper) PersistenceEntityToEntity(pe persistenceentites.CrudPersistenceEntity) model.CrudEntity {
	var gpe persistenceentites.GroceryPersistenceEntity
	switch reflect.TypeOf(pe) {
	case reflect.TypeOf(persistenceentites.GroceryPersistenceEntity{}):
		gpe = pe.(persistenceentites.GroceryPersistenceEntity)
	case reflect.TypeOf(&persistenceentites.GroceryPersistenceEntity{}):
		gpe = *(pe.(*persistenceentites.GroceryPersistenceEntity))
	default:
		gpe = persistenceentites.GroceryPersistenceEntity{}
	}

	flags := []model.Flag{}
	for _, flagPe := range gpe.Flags {
		flag := gm.flagMapper.PersistenceEntityToEntity(flagPe).(*model.Flag)
		flags = append(flags, *flag)
	}

	return &model.Grocery{
		Id:    gpe.Id,
		Name:  gpe.Name,
		Flags: flags,
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
