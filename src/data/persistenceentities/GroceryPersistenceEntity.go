package persistenceentites

type GroceryPersistenceEntity struct {
	Id      string
	Name    string
	FlagIds string
}

func (entity GroceryPersistenceEntity) GetId() string {
	return entity.Id
}
