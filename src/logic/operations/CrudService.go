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

func (service CrudService) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	entites, err := service.repo.ReadAll(target)

	if err != nil {
		return []model.CrudEntity{}, err
	}

	return entites, nil
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
