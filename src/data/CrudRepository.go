package data

import (
	"fmt"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CrudRepository struct {
	db *gorm.DB
}

func NewCrudRepository(db *gorm.DB, entity persistenceentites.CrudPersistenceEntity) CrudRepository {
	db.AutoMigrate(entity)
	return CrudRepository{db: db}
}

func (repo CrudRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	// Create the identifier of the unit
	uuid := uuid.New().String()
	entity.SetId(uuid)

	err := repo.db.Create(entity.ToPersistenceEntity()).Error

	if err != nil {
		return entity.Empty(), fmt.Errorf("something went wrong when adding the unit to the database")
	}

	return entity, nil
}
