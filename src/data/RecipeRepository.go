package data

import (
	"mealwhile/data/mappers"

	"gorm.io/gorm"
)

type RecipeRepository struct {
	db          *gorm.DB
	crudRepo    CrudRepositoryInterface
	crudMappers mappers.CrudMappersInterface
}

/*
func NewRecipeRepository(db *gorm.DB) RecipeRepository {
	db.AutoMigrate(&persistenceentites.RecipePersistenceEntity{})
	crudMappers := mappers.RecipeMapper{}
	crudRepo := NewCrudRepository(db, &persistenceentites.RecipePersistenceEntity{}, crudMappers)
	return RecipeRepository{db: db, crudRepo: crudRepo, crudMappers: crudMappers}
}

func (repo RecipeRepository) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	name := entity.Attributes()["name"]
	// Check if the recipe with the given name already exists
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

func (repo RecipeRepository) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return repo.crudRepo.ReadAll(target)
}

func (repo RecipeRepository) Read(target model.CrudEntity, id string) (model.CrudEntity, error) {
	return repo.crudRepo.Read(target, id)
}

func (repo RecipeRepository) Update(target model.CrudEntity) (model.CrudEntity, error) {
	// Check if the new name clashes with another entity
	recipe := target.(*model.Recipe)
	foundByName, err := repo.FindByName(recipe.Name)

	if err != nil && err.(errors.AppError).Code == 404 {
		// Name does not exist yet
		return repo.crudRepo.Update(target)
	}

	if err != nil {
		// Some sort of db error
		return &model.Recipe{}, err
	}

	if foundByName.GetId() == target.GetId() {
		// The found entity is the entity that is tried to be updated
		return repo.crudRepo.Update(target)
	}

	// Another entity already has the given name
	return &model.Recipe{}, errors.NewEntityAlreadyExists(target, fmt.Sprintf("name %s", recipe.Name))
}

func (repo RecipeRepository) Delete(target model.CrudEntity, id string) error {
	return repo.crudRepo.Delete(target, id)
}

func (repo RecipeRepository) Exists(target model.CrudEntity, id string) (bool, error) {
	return repo.crudRepo.Exists(target, id)
}

func (repo RecipeRepository) FindByName(name string) (model.CrudEntity, error) {
	pe := &persistenceentites.RecipePersistenceEntity{}

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

func (repo RecipeRepository) GroceryReferenced(groceryId string) (bool, error) {
	//recipes, err := repo.ReadAll(&model.Recipe{})

	/*
		if err != nil {
			return false, err
		}



		for _, recipe := range recipes {

				if slices.Contains(recipe.(*model.Recipe).Ingredients, groceryId) {
					return true, nil
				}
		}


	return false, nil
}
*/
