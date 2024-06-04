package data

import (
	"fmt"
	"mealwhile/data/mappers"
	persistenceentites "mealwhile/data/persistenceentities"
	"mealwhile/errors"
	"mealwhile/logic/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GroceryRepository struct {
	db          *gorm.DB
	crudRepo    CrudRepositoryInterface
	crudMappers mappers.CrudMappersInterface
}

func NewGroceryRepository(db *gorm.DB) GroceryRepository {
	db.AutoMigrate(&persistenceentites.GroceryPersistenceEntity{})
	db.Session(&gorm.Session{}).Updates(&persistenceentites.GroceryPersistenceEntity{})
	flagMapper := mappers.FlagMapper{}
	crudMappers := mappers.NewGroceryMapper(&flagMapper)
	crudRepo := NewCrudRepository(db, &persistenceentites.GroceryPersistenceEntity{}, crudMappers, &model.Grocery{})
	return GroceryRepository{db: db, crudRepo: crudRepo, crudMappers: crudMappers}
}

func (repo GroceryRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	name := entity.Attributes()["name"].(string)
	// Check if the grocery with the given name already exists
	_, err := repo.FindByName(name)

	if err == nil {
		return entity.Empty(), errors.NewEntityAlreadyExists(entity, fmt.Sprintf("name %s", name))
	}

	if err != nil && err.(errors.AppError).Code == 404 {
		// Entity not found -> Create the entity
		return repo.crudRepo.Create(entity)
	}

	return entity.Empty(), err
}

func (repo GroceryRepository) ReadAll() ([]model.CrudEntity, error) {
	pes := []persistenceentites.GroceryPersistenceEntity{}
	err := repo.db.Preload(clause.Associations).Model(&persistenceentites.GroceryPersistenceEntity{}).Find(&pes).Error

	if err != nil {
		return nil, errors.NewServerError("Something went wrong retrieving the groceries")
	}

	var result []model.CrudEntity

	for _, gpe := range pes {
		grocery := repo.crudMappers.PersistenceEntityToEntity(gpe)
		result = append(result, grocery)
	}

	return result, nil
}

func (repo GroceryRepository) Read(id string) (model.CrudEntity, error) {
	pe := persistenceentites.GroceryPersistenceEntity{}
	err := repo.db.Preload(clause.Associations).Model(&persistenceentites.GroceryPersistenceEntity{}).Where("id = ?", id).Find(&pe).Error

	if err != nil {
		return nil, errors.NewServerError(fmt.Sprintf("Something went wrong retrieving the grocery with id %s", id))
	}

	result := repo.crudMappers.PersistenceEntityToEntity(pe)

	return result, nil
}

func (repo GroceryRepository) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	// Check if the new name clashes with another entity
	grocery := entity.(*model.Grocery)
	foundByName, err := repo.FindByName(grocery.Name)

	if err != nil && err.(errors.AppError).Code != 404 {
		// Some sort of db error
		return &model.Grocery{}, err
	}

	logrus.Warn("FOUND BY NAME", foundByName)

	if foundByName.GetId() != "" && foundByName.GetId() != entity.GetId() {
		// Another entity already has the given name
		return &model.Grocery{}, errors.NewEntityAlreadyExists(entity, fmt.Sprintf("name %s", grocery.Name))
	}

	pe := repo.crudMappers.EntityToPersistenceEntity(entity)

	// Clear the associations of the grocery
	err = repo.db.Model(&pe).Association("Flags").Replace(pe.(persistenceentites.GroceryPersistenceEntity).Flags)

	if err != nil {
		return nil, errors.NewServerError(fmt.Sprintf("Something went wrong updating the grocery with id %s", pe.GetId()))
	}

	gpe := pe.(persistenceentites.GroceryPersistenceEntity)
	err = repo.db.Save(&gpe).Error

	if err != nil {
		return nil, errors.NewServerError(fmt.Sprintf("Something went wrong updating the grocery with id %s", pe.GetId()))
	}

	return repo.crudMappers.PersistenceEntityToEntity(pe), nil
}

func (repo GroceryRepository) Delete(id string) error {
	// Get the persistence entity, if it exists
	pe, err := repo.ReadPe(id)

	if err != nil {
		return err
	}

	gpe := (*pe).(persistenceentites.GroceryPersistenceEntity)

	// Clear the associations of the grocery
	err = repo.db.Model(&gpe).Association("Flags").Clear()

	if err != nil {
		return errors.NewServerError(fmt.Sprintf("Something went wrong deleting the grocery with id %s", id))
	}

	err = repo.db.Model(&persistenceentites.GroceryPersistenceEntity{}).Delete(&gpe).Error

	if err != nil {
		return errors.NewServerError(fmt.Sprintf("Something went wrong deleting the grocery with id %s", id))
	}

	return nil
}

func (repo GroceryRepository) Exists(id string) (bool, error) {
	return repo.crudRepo.Exists(id)
}

func (repo GroceryRepository) FindByName(name string) (model.CrudEntity, error) {
	pe := &persistenceentites.GroceryPersistenceEntity{}

	err := repo.db.Where("name = ?", name).Find(pe).Error

	if err != nil {
		message := fmt.Sprintf("Something went wrong retrieving the grocery with name %s", name)
		return &model.Grocery{}, errors.NewServerError(message)
	}

	grocery := repo.crudMappers.PersistenceEntityToEntity(*pe)

	if (err != nil && err == gorm.ErrRecordNotFound) || (grocery.(*model.Grocery).Id == "") {
		return &model.Grocery{}, errors.NewEntityNotFound(&model.Grocery{}, fmt.Sprintf("name %s", name))
	} else if err != nil {
		message := fmt.Sprintf("Something went wrong retrieving the grocery with name %s", name)
		return &model.Grocery{}, errors.NewServerError(message)
	}

	return grocery, nil
}

func (repo GroceryRepository) FlagReferenced(flagId string) (bool, error) {
	groceries, err := repo.ReadAll()

	if err != nil {
		return false, err
	}

	for _, grocery := range groceries {
		for _, flag := range grocery.(*model.Grocery).Flags {
			if flag.Id == flagId {
				return true, nil
			}
		}
	}

	return false, nil
}

func (repo GroceryRepository) ReadPe(id string) (*persistenceentites.CrudPersistenceEntity, error) {
	pe := persistenceentites.GroceryPersistenceEntity{}
	err := repo.db.Preload(clause.Associations).Model(&persistenceentites.GroceryPersistenceEntity{}).Where("id = ?", id).Find(&pe).Error

	if err != nil {
		return nil, errors.NewServerError(fmt.Sprintf("Something went wrong retrieving the grocery with id %s", id))
	}

	var result persistenceentites.CrudPersistenceEntity = pe

	return &result, nil
}
