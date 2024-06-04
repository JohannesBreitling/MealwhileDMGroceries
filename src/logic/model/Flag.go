package model

import (
	"fmt"
	"mealwhile/logic/model/requests"
)

type Flag struct {
	Id          string
	Name        string
	Description string
}

func (f Flag) GetId() string {
	return f.Id
}

func (f *Flag) SetId(id string) {
	f.Id = id
}

func (f *Flag) Empty() CrudEntity {
	return &Flag{}
}

func (Flag) EntityName() string {
	return "flag"
}

func (f Flag) Attributes() map[string]interface{} {
	result := make(map[string]interface{})

	result["id"] = f.Id
	result["name"] = f.Name
	result["description"] = f.Description

	return result
}

func (f Flag) String() string {
	str := fmt.Sprintf("{'id': %s, 'name': %s, 'description': %s}", f.Id, f.Name, f.Description)
	return str
}

func (f Flag) FromInterface(arg map[string]interface{}) CrudEntity {
	flag := Flag{}

	id, idOk := arg["id"].(string)
	name, nameOk := arg["name"].(string)
	description, descriptionOk := arg["description"].(string)

	if idOk {
		flag.Id = id
	}

	if nameOk {
		flag.Name = name
	}

	if descriptionOk {
		flag.Description = description
	}

	return &flag
}

func (f Flag) BuildRequest(arg map[string]interface{}) requests.CrudRequest {
	flagRequest := requests.FlagRequest{}

	id, idOk := arg["id"].(string)
	name, nameOk := arg["name"].(string)
	description, descriptionOk := arg["description"].(string)

	if idOk {
		flagRequest.Id = id
	}

	if nameOk {
		flagRequest.Name = name
	}

	if descriptionOk {
		flagRequest.Description = description
	}

	return flagRequest
}
