package operations

import (
	"fmt"
	"mealwhile/data"
	"mealwhile/errors"
	"mealwhile/logic/model"
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

func (ops GroceryOperations) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	// Check if all the flags exist
	grocery := entity.(*model.Grocery)

	for _, flag := range grocery.FlagIds {
		exists, err := ops.flagOperations.Service.repo.Exists(&model.Flag{}, flag)

		if err != nil {
			return &model.Grocery{}, err
		}

		if !exists {
			return &model.Grocery{}, errors.NewEntityNotFound(&model.Flag{}, fmt.Sprintf("id %s", flag))
		}
	}

	return ops.Service.Create(entity)
}

func (ops GroceryOperations) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return ops.Service.ReadAll(&model.Grocery{})
}

func (ops GroceryOperations) Read(target model.CrudEntity, id string) (model.CrudEntity, error) {
	return ops.Service.Read(target, id)
}

func (ops GroceryOperations) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	return ops.Service.Update(entity)
}

func (ops GroceryOperations) Delete(target model.CrudEntity, id string) error {
	return ops.Service.Delete(target, id)
}

func (ops GroceryOperations) FlagReferenced(flagId string) (bool, error) {
	return (ops.Service.repo.(data.GroceryRepositoryInterface)).FlagReferenced(flagId)
}
