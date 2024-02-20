package persistenceentites

type UnitPersistenceEntity struct {
	Id           string
	Name         string
	Abbreviation string
}

func (entity UnitPersistenceEntity) GetId() string {
	return entity.Id
}
