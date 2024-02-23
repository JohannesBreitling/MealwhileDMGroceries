package data

import (
	"mealwhile/data/mappers"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"

	"gorm.io/gorm"
)

type FlagRepository struct {
	db          *gorm.DB
	crudRepo    CrudRepositoryInterface
	crudMappers mappers.CrudMappersInterface
}

func NewFlagRepository(db *gorm.DB) FlagRepository {
	db.AutoMigrate(&persistenceentites.FlagPersistenceEntity{})
	crudMappers := mappers.FlagMapper{}
	crudRepo := NewCrudRepository(db, &persistenceentites.FlagPersistenceEntity{}, crudMappers)
	return FlagRepository{db: db, crudRepo: crudRepo, crudMappers: crudMappers}
}

func (repo FlagRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	return repo.crudRepo.Create(entity)
}

func (repo FlagRepository) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return repo.crudRepo.ReadAll(target)
}

func (repo FlagRepository) Read(target model.CrudEntity, id string) (model.CrudEntity, error) {
	return repo.crudRepo.Read(target, id)
}

func (repo FlagRepository) Update(target model.CrudEntity) (model.CrudEntity, error) {
	return repo.crudRepo.Update(target)
}

func (repo FlagRepository) Delete(target model.CrudEntity, id string) error {
	return repo.crudRepo.Delete(target, id)
}

func (repo FlagRepository) Exists(target model.CrudEntity, id string) (bool, error) {
	return repo.crudRepo.Exists(target, id)
}
