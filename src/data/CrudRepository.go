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
		entity := model.Unit{}.FromInterface(v)
		results = append(results, entity)
	}

	return results, nil
}
