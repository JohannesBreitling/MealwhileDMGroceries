package controller

import (
	"mealwhile/logic/model"

	"github.com/labstack/echo/v4"
)

type CrudControllerInterface interface {
	Create(ctx echo.Context, entity model.CrudEntity) error
	//Get(ctx echo.Context, id string) error
	//GetAll(ctx echo.Context) error
	// Update()
	// Delete()
}
