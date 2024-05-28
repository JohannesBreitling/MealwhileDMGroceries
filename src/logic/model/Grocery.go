package model

import (
	"encoding/json"
	"fmt"
)

type Grocery struct {
	Id      string
	Name    string
	FlagIds []string
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

func flagIdsToString(flags []string) string {
	jsonFlags, _ := json.Marshal(flags)
	return string(jsonFlags)
}

func stringToFlagIds(s string) []string {
	var flagIds []string
	json.Unmarshal([]byte(s), &flagIds)
	return flagIds
}

func (g Grocery) Attributes() map[string]string {
	result := make(map[string]string)

	result["id"] = g.Id
	result["name"] = g.Name
	result["flagIds"] = flagIdsToString(g.FlagIds)

	return result
}

func (g Grocery) String() string {
	str := fmt.Sprintf("{'id': %s, 'name': %s, 'flagIds': %s}", g.Id, g.Name, flagIdsToString(g.FlagIds))
	return str
}

func (Grocery) EntityName() string {
	return "grocery"
}

func (Grocery) FromArguments(args map[string]string) CrudEntity {
	grocery := Grocery{}
	id, idOk := args["id"]
	name, nameOk := args["name"]
	flagIdsString, flagsOk := args["flagIds"]

	if idOk {
		grocery.Id = id
	}

	if nameOk {
		grocery.Name = name
	}

	if flagsOk {
		grocery.FlagIds = stringToFlagIds(flagIdsString)
	}

	return &grocery
}

func (Grocery) FromInterface(args map[string]interface{}) CrudEntity {
	grocery := Grocery{}
	id, idOk := args["id"].(string)
	name, nameOk := args["name"].(string)
	flagIdsString1, flags1Ok := args["flagIds"].([]interface{})
	flagIdsString2, flags2Ok := args["flag_ids"].(string)

	if idOk {
		grocery.Id = id
	}

	if nameOk {
		grocery.Name = name
	}

	var flagIds []string
	if flags1Ok && flagIdsString1 != nil {
		for _, flag := range flagIdsString1 {
			flagIds = append(flagIds, flag.(string))
		}
	}

	if flags2Ok && flagIdsString2 != "" {
		json.Unmarshal([]byte(flagIdsString2), &flagIds)
	}

	grocery.FlagIds = flagIds

	return &grocery
}
