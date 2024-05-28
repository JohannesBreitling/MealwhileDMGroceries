package operations

import (
	"mealwhile/errors"
	"mealwhile/logic/model"
)

type FlagOperations struct {
	Service           CrudService
	groceryOperations *GroceryOperations
}

func NewFlagOperations(service CrudService) FlagOperations {
	return FlagOperations{Service: service}
}

func (o *FlagOperations) SetGroceryOperations(groceryOps *GroceryOperations) {
	o.groceryOperations = groceryOps
}

func (ops FlagOperations) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	return ops.Service.Create(entity)
}

func (ops FlagOperations) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return ops.Service.ReadAll(&model.Flag{})
}

func (ops FlagOperations) Read(target model.CrudEntity, id string) (model.CrudEntity, error) {
	return ops.Service.Read(target, id)
}

func (ops FlagOperations) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	return ops.Service.Update(entity)
}

func (ops FlagOperations) Delete(target model.CrudEntity, id string) error {
	// Check if the flag to be deleted is contained in a grocery
	referenced, err := ops.groceryOperations.FlagReferenced(id)

	if err != nil {
		return err
	}

	if referenced {
		return errors.NewBadRequest("The flag cannot be deleted, as it is used by a grocery")
	}

	return ops.Service.Delete(target, id)
}
