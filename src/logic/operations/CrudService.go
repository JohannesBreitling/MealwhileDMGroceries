package operations

import (
	"fmt"
	"mealwhile/data"
	"mealwhile/errors"
	"mealwhile/logic/model"
)

type CrudService struct {
	repo   data.CrudRepositoryInterface
	target model.CrudEntity
}

func NewCrudService(repo data.CrudRepositoryInterface, target model.CrudEntity) CrudService {
	return CrudService{repo: repo, target: target}
}

func (service CrudService) exists(id string) error {
	exists, err := service.repo.Exists(id)

	if err != nil {
		return err
	}

	if !exists {
		return errors.NewEntityNotFound(service.target, fmt.Sprintf("id %s", id))
	}

	return nil
}

func (service CrudService) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	return service.repo.Create(entity)
}

func (service CrudService) ReadAll() ([]model.CrudEntity, error) {
	return service.repo.ReadAll()
}

func (service CrudService) Read(id string) (model.CrudEntity, error) {
	err := service.exists(id)

	if err != nil {
		return service.target.Empty(), err
	}

	return service.repo.Read(id)
}

func (service CrudService) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	err := service.exists(entity.GetId())

	if err != nil {
		return entity.Empty(), err
	}

	return service.repo.Update(entity)
}

func (service CrudService) Delete(id string) error {
	err := service.exists(id)

	if err != nil {
		return err
	}

	return service.repo.Delete(id)
}
