package controller

import (
	"github.com/labstack/echo/v4"
)

type CrudControllerInterface interface {
	Create(ctx echo.Context) error
	Get(ctx echo.Context, id string) error
	GetAll(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context, id string) error
}
