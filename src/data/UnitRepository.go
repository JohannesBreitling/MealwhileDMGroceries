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
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&persistenceentites.UnitPersistenceEntity{})
	crudMappers := mappers.UnitMapper{}
	crudRepo := NewCrudRepository(db, &persistenceentites.UnitPersistenceEntity{}, crudMappers, &model.Unit{})
	return UnitRepository{db: db, crudRepo: crudRepo, crudMappers: crudMappers}
}

func (repo UnitRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
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

func (repo UnitRepository) ReadAll() ([]model.CrudEntity, error) {
	return repo.crudRepo.ReadAll()
}

func (repo UnitRepository) Read(id string) (model.CrudEntity, error) {
	return repo.crudRepo.Read(id)
}

func (repo UnitRepository) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	// Check if the new name clashes with another entity
	unit := entity.(*model.Unit)
	foundByName, err := repo.FindByName(unit.Name)

	if err != nil && err.(errors.AppError).Code == 404 {
		// Name does not exist yet
		return repo.crudRepo.Update(entity)
	}

	if err != nil {
		// Some sort of db error
		return &model.Unit{}, err
	}

	if foundByName.GetId() == entity.GetId() {
		// The found entity is the entity that is tried to be updated
		return repo.crudRepo.Update(entity)
	}

	// Another entity already has the given name
	return &model.Unit{}, errors.NewEntityAlreadyExists(entity, fmt.Sprintf("name %s", unit.Name))
}

func (repo UnitRepository) Delete(id string) error {
	return repo.crudRepo.Delete(id)
}

func (repo UnitRepository) Exists(id string) (bool, error) {
	return repo.crudRepo.Exists(id)
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

func (repo UnitRepository) ReadPe(id string) (*persistenceentites.CrudPersistenceEntity, error) {
	return repo.crudRepo.ReadPe(id)
}
