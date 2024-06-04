package operations

import (
	"mealwhile/errors"
	"mealwhile/logic/model"
	"mealwhile/logic/model/requests"
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

func (ops FlagOperations) Create(request requests.CrudRequest) (model.CrudEntity, error) {
	// Create the entity from the request
	flagRequest := request.(requests.FlagRequest)
	flag := &model.Flag{Name: flagRequest.Name, Description: flagRequest.Description}

	return ops.Service.Create(flag)
}

func (ops FlagOperations) ReadAll() ([]model.CrudEntity, error) {
	return ops.Service.ReadAll()
}

func (ops FlagOperations) Read(id string) (model.CrudEntity, error) {
	return ops.Service.Read(id)
}

func (ops FlagOperations) Update(request requests.CrudRequest) (model.CrudEntity, error) {
	// Create the entity from the request
	flagRequest := request.(requests.FlagRequest)
	flag := &model.Flag{Id: flagRequest.Id, Name: flagRequest.Name, Description: flagRequest.Description}

	return ops.Service.Update(flag)
}

func (ops FlagOperations) Delete(id string) error {
	// Check if the flag to be deleted is contained in a grocery
	referenced, err := ops.groceryOperations.FlagReferenced(id)

	if err != nil {
		return err
	}

	if referenced {
		return errors.NewBadRequest("The flag cannot be deleted, as it is used by a grocery")
	}

	return ops.Service.Delete(id)
}
