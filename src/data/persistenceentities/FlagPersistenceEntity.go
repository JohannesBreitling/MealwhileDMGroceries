package persistenceentites

type FlagPersistenceEntity struct {
	Id          string
	Name        string
	Description string
}

func (entity FlagPersistenceEntity) GetId() string {
	return entity.Id
}
