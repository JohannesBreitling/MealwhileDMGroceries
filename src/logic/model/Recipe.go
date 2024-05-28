package model

type Recipe struct {
	Id          string
	Name        string
	Description string
	Ingredients []Ingredient
	Steps       []string
}

/*
func (r Recipe) GetId() string {
	return r.Id
}

func (r *Recipe) SetId(id string) {
	r.Id = id
}

func (r *Recipe) Empty() CrudEntity {
	return &Recipe{}
}

func (Recipe) EntityName() string {
	return "recipe"
}

func (r Recipe) Attributes() map[string]string {
	result := make(map[string]string)

	result["id"] = r.Id
	result["name"] = r.Name
	result["description"] = r.Description

	// TODO hier fehelen noch steps und ingredients

	return result
}

func (r Recipe) String() string {
	//TODO str := fmt.Sprintf("{'id': %s, 'name': %s, 'description': %s}", f.Id, f.Name, f.Description)
	// return str
	return ""
}

func (r Recipe) FromArguments(args map[string]string) CrudEntity {
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

func (r Recipe) FromInterface(arg map[string]interface{}) CrudEntity {
	recipe := Recipe{}

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
*/
