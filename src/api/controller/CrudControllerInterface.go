package controller

import (
	"mealwhile/logic/model"

	"github.com/labstack/echo/v4"
)

type CrudControllerInterface interface {
	Create(ctx echo.Context, entity model.CrudEntity) error
	Get(ctx echo.Context, target model.CrudEntity, id string) error
	GetAll(ctx echo.Context, target model.CrudEntity) error
	Update(ctx echo.Context, entity model.CrudEntity) error
	Delete(ctx echo.Context, target model.CrudEntity, id string) error
}
