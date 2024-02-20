package data

import (
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"

	"gorm.io/gorm"
)

type UnitRepository struct {
	db       *gorm.DB
	crudRepo CrudRepositoryInterface
}

func NewUnitRepository(db *gorm.DB, crudRepo CrudRepositoryInterface) UnitRepository {
	db.AutoMigrate(&persistenceentites.UnitPersistenceEntity{})
	return UnitRepository{db: db, crudRepo: crudRepo}
}

func (repo UnitRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	return repo.crudRepo.Create(entity)
}

// Read(entity persistenceentites.CrudPersistenceEntity)
// Update(entity persistenceentites.CrudPersistenceEntity) (persistenceentites.CrudPersistenceEntity, error)
//Delete(entity persistenceentites.CrudPersistenceEntity) error
