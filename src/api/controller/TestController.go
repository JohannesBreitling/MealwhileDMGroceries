package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TestController struct {
}

func NewTestController() TestController {
	return TestController{}
}

func (TestController) Test(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Die Änderungen sind auch live, hat das geklappt?!")
}
