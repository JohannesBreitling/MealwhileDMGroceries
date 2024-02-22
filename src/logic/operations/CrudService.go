package operations

import (
	"fmt"
	"mealwhile/data"
	"mealwhile/logic/model"
)

type CrudService struct {
	repo data.CrudRepositoryInterface
}

func NewCrudService(repo data.CrudRepositoryInterface) CrudService {
	return CrudService{repo: repo}
}

func (service CrudService) exists(entity model.CrudEntity, id string) error {
	exists, err := service.repo.Exists(entity, id)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf(fmt.Sprintf("the entity with id %s does not exist", id))
	}

	return nil
}

func (service CrudService) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	// TODO Advanced: Check if entity (name / abbr) already exists
	return service.repo.Create(entity)
}

func (service CrudService) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return service.repo.ReadAll(target)
}

func (service CrudService) Read(entity model.CrudEntity, id string) (model.CrudEntity, error) {
	return service.repo.Read(entity, id)
}

func (service CrudService) Update(entity model.CrudEntity, id string) (model.CrudEntity, error) {
	err := service.exists(entity, id)

	if err != nil {
		return entity.Empty(), err
	}

	return service.repo.Update(entity, id)
}

func (service CrudService) Delete(target model.CrudEntity, id string) error {
	err := service.exists(target, id)

	if err != nil {
		return err
	}

	return service.repo.Delete(target, id)
}
