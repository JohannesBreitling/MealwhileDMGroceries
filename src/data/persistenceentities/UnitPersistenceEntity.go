package persistenceentites

type UnitPersistenceEntity struct {
	Id           string `gorm:"primaryKey"`
	Name         string
	Abbreviation string
}

func (entity UnitPersistenceEntity) GetId() string {
	return entity.Id
}

func (entity UnitPersistenceEntity) Empty() CrudPersistenceEntity {
	return UnitPersistenceEntity{}
}

func (u UnitPersistenceEntity) FromInterface(arg map[string]interface{}) CrudPersistenceEntity {
	unit := UnitPersistenceEntity{}

	id, idOk := arg["id"].(string)
	name, nameOk := arg["name"].(string)
	abbreviation, abbreviationOk := arg["abbreviation"].(string)

	if idOk {
		unit.Id = id
	}

	if nameOk {
		unit.Name = name
	}

	if abbreviationOk {
		unit.Abbreviation = abbreviation
	}

	return &unit
}
