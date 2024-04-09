package data

import (
	"fmt"
	"mealwhile/data/mappers"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/logic/model"

	"github.com/sirupsen/logrus"
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

	_, err := repo.FindByName(entity.Attributes()["name"])

	if err == nil {
		logrus.Warn("MOIN")
		return &model.Unit{}, fmt.Errorf("unit already existing")
	}

	return repo.crudRepo.Create(entity)
}

func (repo UnitRepository) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return repo.crudRepo.ReadAll(target)
}

func (repo UnitRepository) Read(target model.CrudEntity, id string) (model.CrudEntity, error) {
	return repo.crudRepo.Read(target, id)
}

func (repo UnitRepository) Update(target model.CrudEntity) (model.CrudEntity, error) {
	return repo.crudRepo.Update(target)
}

func (repo UnitRepository) Delete(target model.CrudEntity, id string) error {
	return repo.crudRepo.Delete(target, id)
}

func (repo UnitRepository) Exists(target model.CrudEntity, id string) (bool, error) {
	return repo.crudRepo.Exists(target, id)
}

func (repo UnitRepository) FindByName(name string) (model.CrudEntity, error) {
	unit := &persistenceentites.UnitPersistenceEntity{}

	err := repo.db.Find(unit).Where("name = ?", name).Error

	if err != nil {
		return &model.Unit{}, fmt.Errorf("something went wrong retrieving the entity")
	}

	return repo.crudMappers.PersistenceEntityToEntity(*unit), nil
}

// Update(entity persistenceentites.CrudPersistenceEntity) (persistenceentites.CrudPersistenceEntity, error)
//Delete(entity persistenceentites.CrudPersistenceEntity) error
