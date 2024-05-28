package controller

import (
	"mealwhile/logic/model"
	"mealwhile/logic/operations/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GroceryController struct {
	unitOps               interfaces.UnitOperationsInterface
	unitCrudController    CrudControllerInterface
	flagOps               interfaces.FlagOperationsInterface
	flagCrudController    CrudControllerInterface
	groceryOps            interfaces.GroceryOperationsInterface
	groceryCrudController CrudControllerInterface
}

func NewGroceryController(unitOps interfaces.UnitOperationsInterface, unitCrudController CrudControllerInterface, flagOps interfaces.UnitOperationsInterface, flagCrudController CrudControllerInterface, groceryOps interfaces.GroceryOperationsInterface, groceryCrudController CrudControllerInterface) GroceryController {
	return GroceryController{
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
func (GroceryController) Test(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "The test has been successful!")
}

// TODO Das mit dem Loggen ein bisschen cooler machen

func (ctr GroceryController) CreateUnit(ctx echo.Context) error {
	return ctr.unitCrudController.Create(ctx, &model.Unit{})
}

func (ctr GroceryController) GetUnits(ctx echo.Context) error {
	return ctr.unitCrudController.GetAll(ctx, &model.Unit{})
}

func (ctr GroceryController) GetUnit(ctx echo.Context, id Id) error {
	return ctr.unitCrudController.Get(ctx, &model.Unit{}, id)
}

func (ctr GroceryController) DeleteUnit(ctx echo.Context, id Id) error {
	return ctr.unitCrudController.Delete(ctx, &model.Unit{}, id)
}

func (ctr GroceryController) UpdateUnit(ctx echo.Context) error {
	return ctr.unitCrudController.Update(ctx, &model.Unit{})
}

func (ctr GroceryController) GetFlags(ctx echo.Context) error {
	return ctr.flagCrudController.GetAll(ctx, &model.Flag{})
}

func (ctr GroceryController) CreateFlag(ctx echo.Context) error {
	return ctr.flagCrudController.Create(ctx, &model.Flag{})
}

func (ctr GroceryController) UpdateFlag(ctx echo.Context) error {
	return ctr.flagCrudController.Update(ctx, &model.Flag{})
}

func (ctr GroceryController) DeleteFlag(ctx echo.Context, id Id) error {
	return ctr.flagCrudController.Delete(ctx, &model.Flag{}, id)
}

func (ctr GroceryController) GetFlag(ctx echo.Context, id string) error {
	return ctr.flagCrudController.Get(ctx, &model.Flag{}, id)
}

func (ctr GroceryController) CreateGrocery(ctx echo.Context) error {
	return ctr.groceryCrudController.Create(ctx, &model.Grocery{})
}

func (ctr GroceryController) GetGroceries(ctx echo.Context) error {
	return ctr.groceryCrudController.GetAll(ctx, &model.Grocery{})
}

func (ctr GroceryController) GetGrocery(ctx echo.Context, id Id) error {
	return ctr.groceryCrudController.Get(ctx, &model.Grocery{}, id)
}

func (ctr GroceryController) DeleteGrocery(ctx echo.Context, id Id) error {
	return ctr.groceryCrudController.Delete(ctx, &model.Grocery{}, id)
}

func (ctr GroceryController) UpdateGrocery(ctx echo.Context) error {
	return ctr.groceryCrudController.Update(ctx, &model.Grocery{})
}

func (ctr GroceryController) GetRecipes(ctx echo.Context) error {
	return nil
	//return ctr.flagCrudController.GetAll(ctx, &model.Flag{})
}

func (ctr GroceryController) CreateRecipe(ctx echo.Context) error {
	return nil
	//return ctr.flagCrudController.Create(ctx, &model.Flag{})
}

func (ctr GroceryController) UpdateRecipe(ctx echo.Context) error {
	return nil
	//return ctr.flagCrudController.Update(ctx, &model.Flag{})
}

func (ctr GroceryController) DeleteRecipe(ctx echo.Context, id Id) error {
	return nil
	//return ctr.flagCrudController.Delete(ctx, &model.Flag{}, id)
}

func (ctr GroceryController) GetRecipe(ctx echo.Context, id string) error {
	return nil
	//return ctr.flagCrudController.Get(ctx, &model.Flag{}, id)
}
