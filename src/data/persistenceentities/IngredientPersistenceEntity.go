package persistenceentites

import "fmt"

type IngredientPersistenceEntity struct {
	GroceryId string
	UnitId    string
	Amount    float64
}

func (entity IngredientPersistenceEntity) GetId() string {
	return fmt.Sprintf("%s#%s", entity.GroceryId, entity.UnitId)
}
