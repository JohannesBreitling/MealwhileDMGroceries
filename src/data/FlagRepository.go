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
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&persistenceentites.FlagPersistenceEntity{})
	crudMappers := mappers.FlagMapper{}
	crudRepo := NewCrudRepository(db, &persistenceentites.FlagPersistenceEntity{}, crudMappers, &model.Flag{})
	return FlagRepository{db: db, crudRepo: crudRepo, crudMappers: crudMappers}
}

func (repo FlagRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	name := entity.Attributes()["name"].(string)
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

func (repo FlagRepository) ReadAll() ([]model.CrudEntity, error) {
	return repo.crudRepo.ReadAll()
}

func (repo FlagRepository) Read(id string) (model.CrudEntity, error) {
	return repo.crudRepo.Read(id)
}

func (repo FlagRepository) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	// Check if the new name clashes with another entity
	flag := entity.(*model.Flag)
	foundByName, err := repo.FindByName(flag.Name)

	if err != nil && err.(errors.AppError).Code == 404 {
		// Name does not exist yet
		return repo.crudRepo.Update(entity)
	}

	if err != nil {
		// Some sort of db error
		return &model.Flag{}, err
	}

	if foundByName.GetId() == entity.GetId() {
		// The found entity is the entity that is tried to be updated
		return repo.crudRepo.Update(entity)
	}

	// Another entity already has the given name
	return &model.Flag{}, errors.NewEntityAlreadyExists(entity, fmt.Sprintf("name %s", entity.(*model.Flag).Name))
}

func (repo FlagRepository) Delete(id string) error {
	return repo.crudRepo.Delete(id)
}

func (repo FlagRepository) Exists(id string) (bool, error) {
	return repo.crudRepo.Exists(id)
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

func (repo FlagRepository) ReadPe(id string) (*persistenceentites.CrudPersistenceEntity, error) {
	return repo.crudRepo.ReadPe(id)
}
