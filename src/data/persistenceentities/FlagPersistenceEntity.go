package persistenceentites

type FlagPersistenceEntity struct {
	Id          string `gorm:"primaryKey;"`
	Name        string
	Description string
}

func (entity FlagPersistenceEntity) GetId() string {
	return entity.Id
}

func (entity FlagPersistenceEntity) Empty() CrudPersistenceEntity {
	return FlagPersistenceEntity{}
}

func (f FlagPersistenceEntity) FromInterface(arg map[string]interface{}) CrudPersistenceEntity {
	flag := FlagPersistenceEntity{}

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
