package operations

import (
	"mealwhile/data"
	"mealwhile/logic/model"
	"mealwhile/logic/model/requests"
)

type GroceryOperations struct {
	Service        CrudService
	flagOperations *FlagOperations
}

func NewGroceryOperations(service CrudService) GroceryOperations {
	return GroceryOperations{Service: service}
}

func (o *GroceryOperations) SetFlagOperations(flagOps *FlagOperations) {
	o.flagOperations = flagOps
}

func (ops GroceryOperations) getFlagsFromStrings(flagIds []string) ([]model.Flag, error) {
	// Check if all the flags exist and read them
	var result []model.Flag

	for _, flagId := range flagIds {
		flag, err := ops.flagOperations.Service.repo.Read(flagId)

		if err != nil {
			return nil, err
		}

		result = append(result, *flag.(*model.Flag))
	}

	return result, nil
}

func (ops GroceryOperations) Create(request requests.CrudRequest) (model.CrudEntity, error) {
	// Check if all the flags exist
	groceryRequest := request.(requests.GroceryRequest)

	flags, err := ops.getFlagsFromStrings(groceryRequest.Flags)

	if err != nil {
		return &model.Grocery{}, err
	}

	grocery := model.Grocery{
		Name:  groceryRequest.Name,
		Flags: flags,
	}

	return ops.Service.Create(&grocery)
}

func (ops GroceryOperations) ReadAll() ([]model.CrudEntity, error) {
	return ops.Service.ReadAll()
}

func (ops GroceryOperations) Read(id string) (model.CrudEntity, error) {
	return ops.Service.Read(id)
}

func (ops GroceryOperations) Update(request requests.CrudRequest) (model.CrudEntity, error) {
	// Check if all the flags exist
	groceryRequest := request.(requests.GroceryRequest)

	flags, err := ops.getFlagsFromStrings(groceryRequest.Flags)

	if err != nil {
		return &model.Grocery{}, err
	}

	grocery := model.Grocery{
		Id:    groceryRequest.Id,
		Name:  groceryRequest.Name,
		Flags: flags,
	}

	return ops.Service.Update(&grocery)
}

func (ops GroceryOperations) Delete(id string) error {
	return ops.Service.Delete(id)
}

func (ops GroceryOperations) FlagReferenced(flagId string) (bool, error) {
	return (ops.Service.repo.(data.GroceryRepositoryInterface)).FlagReferenced(flagId)
}
