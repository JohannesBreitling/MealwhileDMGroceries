package model

import (
	"fmt"
	persistenceentites "mealwhile/data/persistenceentities"
)

type Unit struct {
	Id           string
	Name         string
	Abbreviation string
}

func (u Unit) GetId() string {
	return u.Id
}

func (u *Unit) SetId(id string) {
	u.Id = id
}

func (u *Unit) Empty() CrudEntity {
	return &Unit{}
}

func (Unit) EntityName() string {
	return "unit"
}

func (u Unit) String() string {
	str := fmt.Sprintf("{'id': %s, 'name': %s, 'abbreviation': %s}", u.Id, u.Name, u.Abbreviation)
	return str
}

func (u *Unit) ToPersistenceEntity() persistenceentites.CrudPersistenceEntity {
	return persistenceentites.UnitPersistenceEntity{
		Id:           u.Id,
		Name:         u.Name,
		Abbreviation: u.Abbreviation,
	}
}

func (u Unit) FromArguments(args map[string]string) CrudEntity {
	unit := Unit{}

	id, idOk := args["id"]
	name, nameOk := args["name"]
	abbreviation, abbreviationOk := args["abbreviation"]

	if idOk {
		unit.Id = id
	}

	if nameOk {
		unit.Name = name
	}

	if abbreviationOk {
		unit.Abbreviation = abbreviation
	}

	return &unit
}
