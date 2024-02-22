package operations

import "mealwhile/logic/model"

type UnitOperations struct {
	Service CrudService
}

func NewUnitOperations(service CrudService) UnitOperations {
	return UnitOperations{Service: service}
}

func (ops UnitOperations) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	return ops.Service.Create(entity)
}

func (ops UnitOperations) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return ops.Service.ReadAll(&model.Unit{})
}

func (ops UnitOperations) Read(id string) (model.CrudEntity, error) {
	return ops.Service.Read(id)
}

func (ops UnitOperations) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	return ops.Service.Update(entity)
}

func (ops UnitOperations) Delete(id string) error {
	return ops.Service.Delete(id)
}
