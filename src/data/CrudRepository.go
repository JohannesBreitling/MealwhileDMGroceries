package data

import (
	"fmt"
	"mealwhile/data/mappers"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/errors"
	"mealwhile/logic/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CrudRepository struct {
	db          *gorm.DB
	crudMappers mappers.CrudMappersInterface
	Target      model.CrudEntity
}

func NewCrudRepository(db *gorm.DB, entity persistenceentites.CrudPersistenceEntity, crudMappers mappers.CrudMappersInterface, target model.CrudEntity) CrudRepository {
	db.AutoMigrate(entity)
	return CrudRepository{db: db, crudMappers: crudMappers, Target: target}
}

func (repo CrudRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	// Create the identifier of the unit
	uuid := uuid.New().String()
	entity.SetId(uuid)

	pe := repo.crudMappers.EntityToPersistenceEntity(entity)

	err := repo.db.Create(pe).Error

	if err != nil {
		message := fmt.Sprintf("Creation of %s (%s) was not successful", entity.EntityName(), entity.String())
		return entity.Empty(), errors.NewServerError(message)
	}

	err = repo.db.Save(pe).Error

	if err != nil {
		message := fmt.Sprintf("Creation of %s (%s) was not successful", entity.EntityName(), entity.String())
		return entity.Empty(), errors.NewServerError(message)
	}

	return entity, nil
}

func (repo CrudRepository) ReadAll() ([]model.CrudEntity, error) {
	var result []map[string]interface{}

	// Convert the received target to a persistence entity
	petarget := repo.crudMappers.EntityToPersistenceEntity(repo.Target)

	// Find all the results
	err := repo.db.Model(petarget).Find(&result).Error

	if err != nil {
		message := fmt.Sprintf("Retrieval of all entities of type %s was not successful", repo.Target.EntityName())
		return nil, errors.NewServerError(message)
	}

	// Convert the results
	results := []model.CrudEntity{}

	for _, v := range result {
		entity := repo.Target.FromInterface(v)
		results = append(results, entity)
	}

	return results, nil
}

func (repo CrudRepository) Read(id string) (model.CrudEntity, error) {
	pe, err := repo.ReadPe(id)

	if err != nil {
		return nil, err
	}

	return repo.crudMappers.PersistenceEntityToEntity(*pe), nil
	/*
		var result map[string]interface{}

		// Convert the received target to a persistence entity
		petarget := repo.crudMappers.EntityToPersistenceEntity(repo.target)

		// Find the result
		err := repo.db.Model(petarget).Where("id = ?", id).Find(&result).Error

		if err == gorm.ErrRecordNotFound || result == nil {
			return nil, errors.NewEntityNotFound(repo.target, fmt.Sprintf("id %s", repo.target.GetId()))
		}

		if err != nil {
			message := fmt.Sprintf("Retrieval of %s with id %s was not successful", repo.target.EntityName(), id)
			return nil, errors.NewServerError(message)
		}

		// Convert the result
		entity := repo.target.FromInterface(result)

		return entity, nil
	*/
}

func (repo CrudRepository) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	// Check if the entity exists
	found, err := repo.Exists(entity.GetId())

	if err != nil {
		// Some sort of db error has occured
		return entity.Empty(), err
	}

	if !found {
		// The entity that should be updated does not exist
		return entity.Empty(), errors.NewEntityNotFound(entity, fmt.Sprintf("id %s", entity.GetId()))
	}

	// Convert to persistence entity
	pe := repo.crudMappers.EntityToPersistenceEntity(entity)

	err = repo.db.Save(pe).Error

	if err != nil {
		message := fmt.Sprintf("Something went wrong updating the %s (%s) with id %s", entity.EntityName(), entity.String(), entity.GetId())
		return entity.Empty(), errors.NewServerError(message)
	}

	return entity, nil
}

func (repo CrudRepository) Delete(id string) error {
	// Get the persistence entity with the give id

	// Check if a entity with the given id exists
	found, err := repo.Exists(id)

	if err != nil {
		// db error
		return err
	}

	if !found {
		// The entity that should be deleted does not exist
		return errors.NewEntityNotFound(repo.Target, fmt.Sprintf("id %s", id))
	}

	// Get the persistence entity
	pe := repo.crudMappers.EntityToPersistenceEntity(repo.Target)

	foundEntity, err := repo.Read(id)

	if err != nil {
		return err
	}

	err = repo.db.Model(pe).Where("id = ?", id).Delete(foundEntity).Error

	if err != nil {
		return errors.NewServerError("Something went wrong deleting the entity")
	}

	return nil
}

func (repo CrudRepository) Exists(id string) (bool, error) {
	_, err := repo.Read(id)

	if err != nil && err.(errors.AppError).Code == 404 {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (repo CrudRepository) ReadPe(id string) (*persistenceentites.CrudPersistenceEntity, error) {
	var result map[string]interface{}

	// Convert the received target to a persistence entity
	petarget := repo.crudMappers.EntityToPersistenceEntity(repo.Target)

	// Find the result
	err := repo.db.Model(&petarget).Where("id = ?", id).Find(&result).Error

	if err == gorm.ErrRecordNotFound || result == nil {
		return nil, errors.NewEntityNotFound(repo.Target, fmt.Sprintf("id %s", id))
	}

	if err != nil {
		message := fmt.Sprintf("Retrieval of %s with id %s was not successful", repo.Target.EntityName(), id)
		return nil, errors.NewServerError(message)
	}

	// Convert the result
	entity := petarget.FromInterface(result)

	return &entity, nil
}
