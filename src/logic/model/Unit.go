package model

import (
	"fmt"
	"mealwhile/logic/model/requests"
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

func (u Unit) Attributes() map[string]interface{} {
	result := make(map[string]interface{})

	result["id"] = u.Id
	result["name"] = u.Name
	result["abbreviation"] = u.Abbreviation

	return result
}

func (u Unit) FromInterface(arg map[string]interface{}) CrudEntity {
	unit := Unit{}

	id, idOk := arg["id"].(string)
	name, nameOk := arg["name"].(string)
	abbreviation, abbreviationOk := arg["abbreviation"].(string)

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

func (u Unit) BuildRequest(arg map[string]interface{}) requests.CrudRequest {
	unitRequest := requests.UnitRequest{}

	id, idOk := arg["id"]
	name, nameOk := arg["name"]
	abbreviation, abbreviationOk := arg["abbreviation"]

	if idOk {
		unitRequest.Id = id.(string)
	}

	if nameOk {
		unitRequest.Name = name.(string)
	}

	if abbreviationOk {
		unitRequest.Abbreviation = abbreviation.(string)
	}

	return unitRequest
}
