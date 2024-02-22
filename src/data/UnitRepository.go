package data

import (
	"mealwhile/data/mappers"
	persistenceentites "mealwhile/data/persistenceentities"
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
	crudMappers := mappers.UnitMapper{}
	crudRepo := NewCrudRepository(db, &persistenceentites.UnitPersistenceEntity{}, crudMappers)
	return UnitRepository{db: db, crudRepo: crudRepo, crudMappers: crudMappers}
}

func (repo UnitRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	return repo.crudRepo.Create(entity)
}

func (repo UnitRepository) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	tg, err := repo.crudRepo.ReadAll(target)
	return tg, err
}

// Read(entity persistenceentites.CrudPersistenceEntity)
// Update(entity persistenceentites.CrudPersistenceEntity) (persistenceentites.CrudPersistenceEntity, error)
//Delete(entity persistenceentites.CrudPersistenceEntity) error
