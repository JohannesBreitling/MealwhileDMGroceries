package persistenceentites

type RecipePersistenceEntity struct {
	Id          string `gorm:"primaryKey"`
	Name        string
	Description string
	Ingredients []IngredientPersistenceEntity
	Steps       []string
}

func (entity RecipePersistenceEntity) GetId() string {
	return entity.Id
}
