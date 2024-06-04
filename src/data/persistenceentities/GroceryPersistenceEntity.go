package persistenceentites

type GroceryPersistenceEntity struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	Flags []FlagPersistenceEntity `gorm:"many2many;"`
}

func (entity GroceryPersistenceEntity) GetId() string {
	return entity.Id
}

func (entity GroceryPersistenceEntity) Empty() CrudPersistenceEntity {
	return GroceryPersistenceEntity{}
}

func (GroceryPersistenceEntity) FromInterface(args map[string]interface{}) CrudPersistenceEntity {
	grocery := GroceryPersistenceEntity{}
	id, idOk := args["id"].(string)
	name, nameOk := args["name"].(string)
	flagIdsString, flagsOk := args["flags"].([]interface{})

	if idOk {
		grocery.Id = id
	}

	if nameOk {
		grocery.Name = name
	}

	var flags []FlagPersistenceEntity
	if flagsOk && flagIdsString != nil {
		for _, flag := range flagIdsString {
			flags = append(flags, flag.(FlagPersistenceEntity))
		}
	}

	grocery.Flags = flags

	return &grocery
}
