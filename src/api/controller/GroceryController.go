package controller

import (
	"fmt"
	"mealwhile/logic/model"
	"mealwhile/logic/operations/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type GroceryController struct {
	expectedUnitRequest []string
	expectedUnit        []string

	unitOps interfaces.UnitOperationsInterface
}

func NewGroceryController(unitOps interfaces.UnitOperationsInterface) GroceryController {
	return GroceryController{
		expectedUnitRequest: []string{"name", "abbreviation"},
		expectedUnit:        []string{"id", "name", "abbreviation"},
		unitOps:             unitOps,
	}
}

func (GroceryController) validateInput(input map[string]string, expected []string) bool {
	if len(input) != len(expected) {
		return false
	}

	for _, key := range expected {
		if _, ok := input[key]; !ok {
			return false
		}
	}

	return true
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
	var msg string

	// Get the attributes from the request body
	var attributes map[string]string
	err := ctx.Bind(&attributes)

	// Check if the body has the correct format
	msg = "Request body has bad format. It should include only the attributes name and abbreviation"
	if err != nil || !ctr.validateInput(attributes, ctr.expectedUnitRequest) {
		log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}

	unit := model.Unit{Name: attributes["name"], Abbreviation: attributes["abbreviation"]}
	msg = fmt.Sprintf("Create new unit %s (%s)", unit.Name, unit.Abbreviation)
	log.Info(msg)

	// Create the unit
	createdUnit, err := ctr.unitOps.Create(&unit)

	// If there was an error, the error should be returned
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, createdUnit)
}

func (GroceryController) GetUnits(ctx echo.Context) error {
	return nil
}

func (GroceryController) GetUnit(ctx echo.Context, id Id) error {
	return nil
}

func (GroceryController) DeleteUnit(ctx echo.Context, id Id) error {
	return nil
}

func (GroceryController) UpdateUnit(ctx echo.Context, id Id) error {
	return nil
}
