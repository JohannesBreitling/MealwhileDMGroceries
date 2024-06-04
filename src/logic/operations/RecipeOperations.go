package operations

type RecipeOperations struct {
	Service CrudService
}

/*
func NewRecipeOperations(service CrudService) RecipeOperations {
	return RecipeOperations{Service: service}
}

func (ops RecipeOperations) Create(entity model.CrudEntity) (model.CrudEntity, error) {
	return ops.Service.Create(entity)
}

func (ops RecipeOperations) ReadAll(target model.CrudEntity) ([]model.CrudEntity, error) {
	return ops.Service.ReadAll()
}

func (ops RecipeOperations) Read(target model.CrudEntity, id string) (model.CrudEntity, error) {
	return ops.Service.Read(target, id)
}

func (ops RecipeOperations) Update(entity model.CrudEntity) (model.CrudEntity, error) {
	return ops.Service.Update(entity)
}

func (ops RecipeOperations) Delete(target model.CrudEntity, id string) error {
	return ops.Service.Delete(target, id)
}
*/
