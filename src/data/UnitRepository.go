package data

import (
	"fmt"
	"mealwhile/data/mappers"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/errors"
	"mealwhile/logic/model"

	"gorm.io/gorm"
)

type UnitRepository struct {
	db          *gorm.DB
	crudRepo    CrudRepositoryInterface
	crudMappers mappers.CrudMappersInterface
}

func NewUnitRepository(db *gorm.DB) UnitRepository {
	db.AutoMigrate(&persistenceentites.UnitPersistenceEntity{})
	crudMappers := mappers.UnitMapper{}
	crudRepo := NewCrudRepository(db, &persistenceentites.UnitPersistenceEntity{}, crudMappers)
	return UnitRepository{db: db, crudRepo: crudRepo, crudMappers: crudMappers}
}

func (repo UnitRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	name := entity.Attributes()["name"]
	// Check if the unit with the given name already exists
	_, err := repo.FindByName(name)

	if err == nil {
		return entity.Empty(), errors.NewEntityAlreadyExists(entity, fmt.Sprintf("name %s", name))
	}

	if err != nil && err.(errors.AppError).Code == 404 {
		// Entity not found -> Create the entity
		return repo.crudRepo.Create(entity)
	}

	return entity.Empty(), err
}

func (repo UnitRepository) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return repo.crudRepo.ReadAll(target)
}

func (repo UnitRepository) Read(target model.CrudEntity, id string) (model.CrudEntity, error) {
	return repo.crudRepo.Read(target, id)
}

func (repo UnitRepository) Update(target model.CrudEntity) (model.CrudEntity, error) {
	// Check if the new name clashes with another entity
	unit := target.(*model.Unit)
	foundByName, err := repo.FindByName(unit.Name)

	if err != nil && err.(errors.AppError).Code == 404 {
		// Name does not exist yet
		return repo.crudRepo.Update(target)
	}

	if err != nil {
		// Some sort of db error
		return &model.Unit{}, err
	}

	if foundByName.GetId() == target.GetId() {
		// The found entity is the entity that is tried to be updated
		return repo.crudRepo.Update(target)
	}

	// Another entity already has the given name
	return &model.Unit{}, errors.NewEntityAlreadyExists(target, fmt.Sprintf("name %s", unit.Name))
}

func (repo UnitRepository) Delete(target model.CrudEntity, id string) error {
	return repo.crudRepo.Delete(target, id)
}

func (repo UnitRepository) Exists(target model.CrudEntity, id string) (bool, error) {
	return repo.crudRepo.Exists(target, id)
}

func (repo UnitRepository) FindByName(name string) (model.CrudEntity, error) {
	pe := &persistenceentites.UnitPersistenceEntity{}

	err := repo.db.Where("name = ?", name).Find(pe).Error

	if err != nil {
		return &model.Unit{}, errors.NewServerError(fmt.Sprintf("Something went wrong retrieving the unit with name %s", name))
	}

	unit := repo.crudMappers.PersistenceEntityToEntity(*pe)

	if (err != nil && err == gorm.ErrRecordNotFound) || (*unit.(*model.Unit) == model.Unit{}) {
		return &model.Unit{}, errors.NewEntityNotFound(&model.Unit{}, fmt.Sprintf("name %s", name))
	} else if err != nil {
		message := fmt.Sprintf("Something went wrong retrieving the unit with name %s", name)
		return &model.Unit{}, errors.NewServerError(message)
	}

	return unit, nil
}
