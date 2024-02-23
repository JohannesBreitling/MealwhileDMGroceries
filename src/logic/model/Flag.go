package model

import (
	"fmt"
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

func (f Flag) String() string {
	str := fmt.Sprintf("{'id': %s, 'name': %s, 'description': %s}", f.Id, f.Name, f.Description)
	return str
}

func (f Flag) FromArguments(args map[string]string) CrudEntity {
	flag := Flag{}

	id, idOk := args["id"]
	name, nameOk := args["name"]
	description, descriptionOk := args["description"]

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
