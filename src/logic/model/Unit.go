package model

import persistenceentites "mealwhile/data/persistenceentities"

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

func (u Unit) Empty() CrudEntity {
	return &Unit{}
}

func (u *Unit) ToPersistenceEntity() persistenceentites.CrudPersistenceEntity {
	return persistenceentites.UnitPersistenceEntity{
		Id:           u.Id,
		Name:         u.Name,
		Abbreviation: u.Abbreviation,
	}
}
