package data

import (
	"fmt"
	"mealwhile/data/mappers"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/errors"
	"mealwhile/logic/model"

	"gorm.io/gorm"
)

type FlagRepository struct {
	db          *gorm.DB
	crudRepo    CrudRepositoryInterface
	crudMappers mappers.CrudMappersInterface
}

func NewFlagRepository(db *gorm.DB) FlagRepository {
	db.AutoMigrate(&persistenceentites.FlagPersistenceEntity{})
	crudMappers := mappers.FlagMapper{}
	crudRepo := NewCrudRepository(db, &persistenceentites.FlagPersistenceEntity{}, crudMappers)
	return FlagRepository{db: db, crudRepo: crudRepo, crudMappers: crudMappers}
}

func (repo FlagRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
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

func (repo FlagRepository) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return repo.crudRepo.ReadAll(target)
}

func (repo FlagRepository) Read(target model.CrudEntity, id string) (model.CrudEntity, error) {
	return repo.crudRepo.Read(target, id)
}

func (repo FlagRepository) Update(target model.CrudEntity) (model.CrudEntity, error) {
	// Check if the new name clashes with another entity
	flag := target.(*model.Flag)
	foundByName, err := repo.FindByName(flag.Name)

	if err != nil && err.(errors.AppError).Code == 404 {
		// Name does not exist yet
		return repo.crudRepo.Update(target)
	}

	if err != nil {
		// Some sort of db error
		return &model.Flag{}, err
	}

	if foundByName.GetId() == target.GetId() {
		// The found entity is the entity that is tried to be updated
		return repo.crudRepo.Update(target)
	}

	// Another entity already has the given name
	return &model.Flag{}, errors.NewEntityAlreadyExists(target, fmt.Sprintf("name %s", target.(*model.Flag).Name))
}

func (repo FlagRepository) Delete(target model.CrudEntity, id string) error {
	return repo.crudRepo.Delete(target, id)
}

func (repo FlagRepository) Exists(target model.CrudEntity, id string) (bool, error) {
	return repo.crudRepo.Exists(target, id)
}

func (repo FlagRepository) FindByName(name string) (model.CrudEntity, error) {
	pe := &persistenceentites.FlagPersistenceEntity{}

	err := repo.db.Where("name = ?", name).Find(pe).Error

	if err != nil {
		return &model.Unit{}, errors.NewServerError(fmt.Sprintf("Something went wrong retrieving the flag with name %s", name))
	}

	flag := repo.crudMappers.PersistenceEntityToEntity(*pe)

	if (err != nil && err == gorm.ErrRecordNotFound) || (*flag.(*model.Flag) == model.Flag{}) {
		return &model.Flag{}, errors.NewEntityNotFound(&model.Flag{}, fmt.Sprintf("name %s", name))
	} else if err != nil {
		message := fmt.Sprintf("Something went wrong retrieving the flag with name %s", name)
		return &model.Flag{}, errors.NewServerError(message)
	}

	return flag, nil
}
