package model

import (
	"fmt"
	"mealwhile/logic/model/requests"
)

type Grocery struct {
	Id    string
	Name  string
	Flags []Flag
}

func (g Grocery) GetId() string {
	return g.Id
}

func (g *Grocery) SetId(id string) {
	g.Id = id
}

func (Grocery) Empty() CrudEntity {
	return &Grocery{}
}

func (g Grocery) Attributes() map[string]interface{} {
	result := make(map[string]interface{})

	result["id"] = g.Id
	result["name"] = g.Name
	result["flags"] = g.Flags

	return result
}

func (g Grocery) String() string {
	str := fmt.Sprintf("{'id': %s, 'name': %s, 'flags': %s}", g.Id, g.Name, flagsToString(g.Flags))
	return str
}

func flagsToString(flags []Flag) string {
	var result string = ""

	for _, flag := range flags {
		result += flag.String()
	}

	return result
}

func (Grocery) EntityName() string {
	return "grocery"
}

func (Grocery) FromInterface(args map[string]interface{}) CrudEntity {
	grocery := Grocery{}
	id, idOk := args["id"].(string)
	name, nameOk := args["name"].(string)
	flagIdsString, flagsOk := args["flags"].([]interface{})

	if idOk {
		grocery.Id = id
	}

	if nameOk {
		grocery.Name = name
	}

	var flags []Flag
	if flagsOk && flagIdsString != nil {
		for _, flag := range flagIdsString {
			flags = append(flags, flag.(Flag))
		}
	}

	grocery.Flags = flags

	return &grocery
}

func (Grocery) BuildRequest(args map[string]interface{}) requests.CrudRequest {
	groceryRequest := requests.GroceryRequest{}
	id, idOk := args["id"].(string)
	name, nameOk := args["name"].(string)
	flagIdsString, flagsOk := args["flagIds"].([]interface{})

	if idOk {
		groceryRequest.Id = id
	}

	if nameOk {
		groceryRequest.Name = name
	}

	var flags []string
	if flagsOk && flagIdsString != nil {
		for _, flag := range flagIdsString {
			flags = append(flags, flag.(string))
		}
	}

	groceryRequest.Flags = flags

	return groceryRequest
}
