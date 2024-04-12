package data

import (
	"fmt"
	"mealwhile/data/mappers"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CrudRepository struct {
	db          *gorm.DB
	crudMappers mappers.CrudMappersInterface
}

func NewCrudRepository(db *gorm.DB, entity persistenceentites.CrudPersistenceEntity, crudMappers mappers.CrudMappersInterface) CrudRepository {
	db.AutoMigrate(entity)
	return CrudRepository{db: db, crudMappers: crudMappers}
}

func (repo CrudRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	// Create the identifier of the unit
	uuid := uuid.New().String()
	entity.SetId(uuid)

	err := repo.db.Create(repo.crudMappers.EntityToPersistenceEntity(entity)).Error

	if err != nil {
		message := fmt.Sprintf("Creation of %s (%s) was not successful", entity.EntityName(), entity.String())
		return entity.Empty(), NewDBRepositoryError(message)
	}

	return entity, nil
}

func (repo CrudRepository) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	var result []map[string]interface{}

	// Convert the received target to a persistence entity
	petarget := repo.crudMappers.EntityToPersistenceEntity(target)

	// Find all the results
	err := repo.db.Model(petarget).Find(&result).Error

	if err != nil {
		message := fmt.Sprintf("Retrieval of all entities of type %s was not successful", target.EntityName())
		return nil, NewDBRepositoryError(message)
	}

	// Convert the results
	results := []model.CrudEntity{}

	for _, v := range result {
		entity := target.FromInterface(v)
		results = append(results, entity)
	}

	return results, nil
}

func (repo CrudRepository) Read(target model.CrudEntity, id string) (model.CrudEntity, error) {
	var result map[string]interface{}

	// Convert the received target to a persistence entity
	petarget := repo.crudMappers.EntityToPersistenceEntity(target)

	// Find the result
	err := repo.db.Model(petarget).Where("id = ?", id).Find(&result).Error

	if err == gorm.ErrRecordNotFound || len(result) == 0 {
		return nil, NewNotFoundRepositoryError(target, fmt.Sprintf("id %s", id))
	}

	if err != nil {
		message := fmt.Sprintf("Retrieval of %s with id %s was not successful", target.EntityName(), id)
		return nil, NewDBRepositoryError(message)
	}

	// Convert the result
	entity := target.FromInterface(result)

	return entity, nil
}

func (repo CrudRepository) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	// Check if the entity exists
	found, err := repo.Exists(entity, entity.GetId())

	if err != nil {
		// Some sort of db error has occured
		return entity.Empty(), err
	}

	if !found {
		// The entity that should be updated does not exist
		return entity.Empty(), NewNotFoundRepositoryError(entity, fmt.Sprintf("id %s", entity.GetId()))
	}

	// Convert to persistence entity
	pe := repo.crudMappers.EntityToPersistenceEntity(entity)

	err = repo.db.Save(pe).Error

	if err != nil {
		message := fmt.Sprintf("Something went wrong updating the %s (%s) with id %s", entity.EntityName(), entity.String(), entity.GetId())
		return entity.Empty(), NewDBRepositoryError(message)
	}

	return entity, nil
}

func (repo CrudRepository) Delete(entity model.CrudEntity, id string) error {
	// Check if a entity with the given id exists
	found, err := repo.Exists(entity, id)

	if err != nil {
		// db error
		return err
	}

	if !found {
		// The entity that should be deleted does not exist
		return NewNotFoundRepositoryError(entity, fmt.Sprintf("id %s", id))
	}

	// Get the persistence entity
	pe := repo.crudMappers.EntityToPersistenceEntity(entity)

	foundEntity, err := repo.Read(entity, id)

	if err != nil {
		return fmt.Errorf("something went wrong deleting the entity")
	}

	err = repo.db.Model(pe).Where("id = ?", id).Delete(foundEntity).Error

	if err != nil {
		return fmt.Errorf("something went wrong deleting the entity")
	}

	return nil
}

func (repo CrudRepository) Exists(entity model.CrudEntity, id string) (bool, error) {
	_, err := repo.Read(entity, id)

	if err != nil && err.(RepositoryError).Code == 1 {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
