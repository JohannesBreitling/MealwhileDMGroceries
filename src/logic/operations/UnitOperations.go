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

func (ops UnitOperations) Read(entity model.CrudEntity) (model.CrudEntity, error) {
	return ops.Service.Read(entity)
}

func (ops UnitOperations) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	return ops.Service.Update(entity)
}

func (ops UnitOperations) Delete(entity model.CrudEntity) error {
	return ops.Service.Delete(entity)
}
