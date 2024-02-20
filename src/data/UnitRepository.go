package data

import (
	"fmt"
	"mealwhile/data/mappers"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UnitRepository struct {
	db *gorm.DB
}

func NewUnitRepository(db *gorm.DB) UnitRepository {
	db.AutoMigrate(&persistenceentites.UnitPersistenceEntity{})
	return UnitRepository{db: db}
}

func (repo UnitRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	var msg string

	// Convert the entity to a unit
	unit, ok := entity.(model.Unit)

	// Check if the conversion was sucessful
	if !ok {
		msg = "the conversion of entity to unit was not successful"
		log.Error(msg)
		return model.Unit{}, fmt.Errorf(msg)
	}

	// Create the identifier of the unit
	uuid := uuid.New().String()
	unit.Id = uuid

	err := repo.db.Create(mappers.UnitToUnitPersistenceEntity(unit)).Error

	if err != nil {
		return model.Unit{}, fmt.Errorf("something went wrong when adding the unit to the database")
	}

	return unit, nil
}

// Read(entity persistenceentites.CrudPersistenceEntity)
// Update(entity persistenceentites.CrudPersistenceEntity) (persistenceentites.CrudPersistenceEntity, error)
//Delete(entity persistenceentites.CrudPersistenceEntity) error
