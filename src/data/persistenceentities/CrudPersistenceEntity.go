package persistenceentites

type CrudPersistenceEntity interface {
	GetId() string
	Empty() CrudPersistenceEntity
	FromInterface(arg map[string]interface{}) CrudPersistenceEntity
}
