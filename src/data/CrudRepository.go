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
		return entity.Empty(), fmt.Errorf("something went wrong when adding the unit to the database")
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
		return nil, fmt.Errorf("something went wrong retrieving the entities")
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

	if err != nil {
		return nil, fmt.Errorf("something went wrong retrieving the entitiy")
	}

	// Convert the result
	entity := target.FromInterface(result)

	return entity, nil
}

// TODO ID BEIM UPDATE RAUS
func (repo CrudRepository) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	// Get the persistence entity
	pe := repo.crudMappers.EntityToPersistenceEntity(entity)

	err := repo.db.Save(pe).Error

	if err != nil {
		return entity.Empty(), fmt.Errorf("something went wrong updating the entity")
	}

	return entity, nil
}

func (repo CrudRepository) Delete(entity model.CrudEntity, id string) error {
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
	// Get the persistence entity
	pe := repo.crudMappers.EntityToPersistenceEntity(entity)

	var result map[string]interface{}

	err := repo.db.Model(pe).Where("id = ?", id).Find(&result).Error

	if err != nil && err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("something went wrong updating the entity")
	}

	return true, nil
}
