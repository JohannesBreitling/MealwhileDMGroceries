package controller

import (
	"mealwhile/logic/operations/interfaces"

	"github.com/labstack/echo/v4"
)

type DMRecipeController struct {
	unitOps               interfaces.UnitOperationsInterface
	unitCrudController    CrudControllerInterface
	flagOps               interfaces.FlagOperationsInterface
	flagCrudController    CrudControllerInterface
	groceryOps            interfaces.GroceryOperationsInterface
	groceryCrudController CrudControllerInterface
}

func NewGroceryController(unitOps interfaces.UnitOperationsInterface, unitCrudController CrudControllerInterface, flagOps interfaces.UnitOperationsInterface, flagCrudController CrudControllerInterface, groceryOps interfaces.GroceryOperationsInterface, groceryCrudController CrudControllerInterface) DMRecipeController {
	return DMRecipeController{
		unitOps:               unitOps,
		unitCrudController:    unitCrudController,
		flagOps:               flagOps,
		flagCrudController:    flagCrudController,
		groceryOps:            groceryOps,
		groceryCrudController: groceryCrudController,
	}
}

// --------------------
// Implement handlers for the routes
// --------------------

func (ctr DMRecipeController) CreateUnit(ctx echo.Context) error {
	return ctr.unitCrudController.Create(ctx)
}

func (ctr DMRecipeController) GetUnits(ctx echo.Context) error {
	return ctr.unitCrudController.GetAll(ctx)
}

func (ctr DMRecipeController) GetUnit(ctx echo.Context, id Id) error {
	return ctr.unitCrudController.Get(ctx, id)
}

func (ctr DMRecipeController) DeleteUnit(ctx echo.Context, id Id) error {
	return ctr.unitCrudController.Delete(ctx, id)
}

func (ctr DMRecipeController) UpdateUnit(ctx echo.Context) error {
	return ctr.unitCrudController.Update(ctx)
}

func (ctr DMRecipeController) GetFlags(ctx echo.Context) error {
	return ctr.flagCrudController.GetAll(ctx)
}

func (ctr DMRecipeController) CreateFlag(ctx echo.Context) error {
	return ctr.flagCrudController.Create(ctx)
}

func (ctr DMRecipeController) UpdateFlag(ctx echo.Context) error {
	return ctr.flagCrudController.Update(ctx)
}

func (ctr DMRecipeController) DeleteFlag(ctx echo.Context, id Id) error {
	return ctr.flagCrudController.Delete(ctx, id)
}

func (ctr DMRecipeController) GetFlag(ctx echo.Context, id string) error {
	return ctr.flagCrudController.Get(ctx, id)
}

func (ctr DMRecipeController) CreateGrocery(ctx echo.Context) error {
	return ctr.groceryCrudController.Create(ctx)
}

func (ctr DMRecipeController) GetGroceries(ctx echo.Context) error {
	return ctr.groceryCrudController.GetAll(ctx)
}

func (ctr DMRecipeController) GetGrocery(ctx echo.Context, id Id) error {
	return ctr.groceryCrudController.Get(ctx, id)
}

func (ctr DMRecipeController) DeleteGrocery(ctx echo.Context, id Id) error {
	return ctr.groceryCrudController.Delete(ctx, id)
}

func (ctr DMRecipeController) UpdateGrocery(ctx echo.Context) error {
	return ctr.groceryCrudController.Update(ctx)
}

func (ctr DMRecipeController) GetRecipes(ctx echo.Context) error {
	return nil
	//return ctr.flagCrudController.GetAll(ctx, &model.Flag{})
}

func (ctr DMRecipeController) CreateRecipe(ctx echo.Context) error {
	return nil
	//return ctr.flagCrudController.Create(ctx, &model.Flag{})
}

func (ctr DMRecipeController) UpdateRecipe(ctx echo.Context) error {
	return nil
	//return ctr.flagCrudController.Update(ctx, &model.Flag{})
}

func (ctr DMRecipeController) DeleteRecipe(ctx echo.Context, id Id) error {
	return nil
	//return ctr.flagCrudController.Delete(ctx, &model.Flag{}, id)
}

func (ctr DMRecipeController) GetRecipe(ctx echo.Context, id string) error {
	return nil
	//return ctr.flagCrudController.Get(ctx, &model.Flag{}, id)
}
