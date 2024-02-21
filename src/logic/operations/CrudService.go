package operations

import (
	"mealwhile/data"
	"mealwhile/logic/model"
)

type CrudService struct {
	repo data.CrudRepositoryInterface
}

func NewCrudService(repo data.CrudRepositoryInterface) CrudService {
	return CrudService{repo: repo}
}

func (service CrudService) Create(entity model.CrudEntity) (model.CrudEntity, error) {

	// TODO Advanced: Check if entity (name / abbr) already exists

	createdEntity, err := service.repo.Create(entity)

	if err != nil {
		return createdEntity, err
	}

	return createdEntity, nil
}

func (CrudService) ReadAll() ([]model.CrudEntity, error) {
	// TODO implement
	return nil, nil
}

func (CrudService) Read(id string) (model.CrudEntity, error) {
	// TODO implement
	return nil, nil
}

func (CrudService) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	// TODO implement
	return nil, nil
}

func (CrudService) Delete(id string) error {
	// TODO implement
	return nil
}
