package controller

import (
	"mealwhile/logic/model"
	"mealwhile/logic/operations/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GroceryController struct {
	unitOps            interfaces.UnitOperationsInterface
	unitCrudController CrudControllerInterface
	flagOps            interfaces.FlagOperationsInterface
	flagCrudController CrudControllerInterface
}

func NewGroceryController(unitOps interfaces.UnitOperationsInterface, unitCrudController CrudControllerInterface, flagOps interfaces.UnitOperationsInterface, flagCrudController CrudControllerInterface) GroceryController {
	return GroceryController{
		unitOps:            unitOps,
		unitCrudController: unitCrudController,
		flagOps:            flagOps,
		flagCrudController: flagCrudController,
	}
}

// --------------------
// Implement handlers for the routes
// --------------------
func (GroceryController) Test(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "The test has been successful!")
}

// TODO Das mit dem Loggen ein bisschen cooler machen
// TODO Error handling nochmal überdenken

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
