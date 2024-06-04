package operations

import (
	"mealwhile/logic/model"
	"mealwhile/logic/model/requests"
)

type UnitOperations struct {
	Service CrudService
}

func NewUnitOperations(service CrudService) UnitOperations {
	return UnitOperations{Service: service}
}

func (ops UnitOperations) Create(request requests.CrudRequest) (model.CrudEntity, error) {
	// Build the entity from the request
	unitRequest := request.(requests.UnitRequest)
	entity := &model.Unit{Name: unitRequest.Name, Abbreviation: unitRequest.Abbreviation}

	return ops.Service.Create(entity)
}

func (ops UnitOperations) ReadAll() ([]model.CrudEntity, error) {
	return ops.Service.ReadAll()
}

func (ops UnitOperations) Read(id string) (model.CrudEntity, error) {
	return ops.Service.Read(id)
}

func (ops UnitOperations) Update(request requests.CrudRequest) (model.CrudEntity, error) {
	// Build the entity from the request
	unitRequest := request.(requests.UnitRequest)
	entity := &model.Unit{Id: unitRequest.Id, Name: unitRequest.Name, Abbreviation: unitRequest.Abbreviation}

	return ops.Service.Update(entity)
}

func (ops UnitOperations) Delete(id string) error {
	return ops.Service.Delete(id)
}
