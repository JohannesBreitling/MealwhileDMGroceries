package operations

import (
	"mealwhile/logic/model"
)

type FlagOperations struct {
	Service CrudService
}

func NewFlagOperations(service CrudService) FlagOperations {
	return FlagOperations{Service: service}
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
	return ops.Service.Delete(target, id)
}
