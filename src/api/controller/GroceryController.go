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
}

func NewGroceryController(unitOps interfaces.UnitOperationsInterface, unitCrudController CrudControllerInterface) GroceryController {
	return GroceryController{
		unitOps:            unitOps,
		unitCrudController: unitCrudController,
	}
}

//func (GroceryController) validateInput(input map[string]string, expected []string) bool {
//	if len(input) != len(expected) {
//		return false
//	}
//
//	for _, key := range expected {
//		if _, ok := input[key]; !ok {
//			return false
//		}
//	}
//
// 	return true
// }

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
	units, err := ctr.unitOps.ReadAll()

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, units)
}

func (ctr GroceryController) GetUnit(ctx echo.Context, id Id) error {
	//units, err := ctr.unitOps.Read(id)

	//if err != nil {
	//	return err
	// }

	// return ctx.JSON(http.StatusOK, units)

	return ctx.JSON(http.StatusNotImplemented, "Not implemented")
}

func (ctr GroceryController) DeleteUnit(ctx echo.Context, id Id) error {
	return nil
}

func (ctr GroceryController) UpdateUnit(ctx echo.Context, id Id) error {
	return nil
}
