package data

import (
	"fmt"
	"mealwhile/data/mappers"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/errors"
	"mealwhile/logic/model"
	"slices"

	"gorm.io/gorm"
)

type GroceryRepository struct {
	db          *gorm.DB
	crudRepo    CrudRepositoryInterface
	crudMappers mappers.CrudMappersInterface
}

func NewGroceryRepository(db *gorm.DB) GroceryRepository {
	db.AutoMigrate(&persistenceentites.GroceryPersistenceEntity{})
	crudMappers := mappers.GroceryMapper{}
	crudRepo := NewCrudRepository(db, &persistenceentites.GroceryPersistenceEntity{}, crudMappers)
	return GroceryRepository{db: db, crudRepo: crudRepo, crudMappers: crudMappers}
}

func (repo GroceryRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	name := entity.Attributes()["name"]
	// Check if the grocery with the given name already exists
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

func (repo GroceryRepository) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return repo.crudRepo.ReadAll(target)
}

func (repo GroceryRepository) Read(target model.CrudEntity, id string) (model.CrudEntity, error) {
	return repo.crudRepo.Read(target, id)
}

func (repo GroceryRepository) Update(target model.CrudEntity) (model.CrudEntity, error) {
	// Check if the new name clashes with another entity
	grocery := target.(*model.Grocery)
	foundByName, err := repo.FindByName(grocery.Name)

	if err != nil && err.(errors.AppError).Code == 404 {
		// Name does not exist yet
		return repo.crudRepo.Update(target)
	}

	if err != nil {
		// Some sort of db error
		return &model.Grocery{}, err
	}

	if foundByName.GetId() == target.GetId() {
		// The found entity is the entity that is tried to be updated
		return repo.crudRepo.Update(target)
	}

	// Another entity already has the given name
	return &model.Grocery{}, errors.NewEntityAlreadyExists(target, fmt.Sprintf("name %s", grocery.Name))
}

func (repo GroceryRepository) Delete(target model.CrudEntity, id string) error {
	return repo.crudRepo.Delete(target, id)
}

func (repo GroceryRepository) Exists(target model.CrudEntity, id string) (bool, error) {
	return repo.crudRepo.Exists(target, id)
}

func (repo GroceryRepository) FindByName(name string) (model.CrudEntity, error) {
	pe := &persistenceentites.GroceryPersistenceEntity{}

	err := repo.db.Where("name = ?", name).Find(pe).Error

	if err != nil {
		message := fmt.Sprintf("Something went wrong retrieving the grocery with name %s", name)
		return &model.Grocery{}, errors.NewServerError(message)
	}

	grocery := repo.crudMappers.PersistenceEntityToEntity(*pe)

	if (err != nil && err == gorm.ErrRecordNotFound) || (grocery.(*model.Grocery).Id == "") {
		return &model.Grocery{}, errors.NewEntityNotFound(&model.Grocery{}, fmt.Sprintf("name %s", name))
	} else if err != nil {
		message := fmt.Sprintf("Something went wrong retrieving the grocery with name %s", name)
		return &model.Grocery{}, errors.NewServerError(message)
	}

	return grocery, nil
}

func (repo GroceryRepository) FlagReferenced(flagId string) (bool, error) {
	groceries, err := repo.ReadAll(&model.Grocery{})

	if err != nil {
		return false, err
	}

	for _, grocery := range groceries {
		if slices.Contains(grocery.(*model.Grocery).FlagIds, flagId) {
			return true, nil
		}
	}

	return false, nil
}
