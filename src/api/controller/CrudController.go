package controller

import (
	"fmt"
	"mealwhile/errors"
	"mealwhile/logic/model"
	"mealwhile/logic/operations/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ExpectedCrudArguments struct {
	Create []string
	Update []string
}

func BuildAttributeString(attributes []string) string {
	attributeList := ""
	separator := ""

	for _, s := range attributes {
		attributeString := fmt.Sprintf("'%s': 'value of %s'", s, s)
		attributeList += separator
		attributeList += attributeString
		separator = ", "
	}

	return fmt.Sprintf("{%s}", attributeList)
}

type CrudController struct {
	args   ExpectedCrudArguments
	ops    interfaces.CrudServiceInterface
	target model.CrudEntity
}

func NewCrudController(ops interfaces.CrudServiceInterface, args ExpectedCrudArguments, target model.CrudEntity) CrudController {
	return CrudController{ops: ops, args: args, target: target}
}

func (CrudController) validateInput(input map[string]interface{}, expected []string) bool {
	if len(input) != len(expected) {
		return false
	}

	for _, key := range expected {
		if _, ok := input[key]; !ok || input[key] == "" {
			return false
		}
	}

	return true
}

func (ctr CrudController) Create(ctx echo.Context) error {
	var msg string

	// Get the attributes from the request body
	var attributes map[string]interface{}
	err := ctx.Bind(&attributes)

	if err != nil {
		e := errors.NewServerError("Receiving the request data was not successful").(errors.AppError)
		return ctx.JSON(e.Code, e.Message)
	}

	// Check if the body has the correct format
	valid := ctr.validateInput(attributes, ctr.args.Create)
	msg = fmt.Sprintf("Request body has bad format. It should have the following format: %s", BuildAttributeString(ctr.args.Create))
	if !valid {
		e := errors.NewBadRequest(msg)
		return ctx.JSON(e.(errors.AppError).Code, e.(errors.AppError).Message)
	}

	request := ctr.target.BuildRequest(attributes)

	// Create the entity
	createdEntity, err := ctr.ops.Create(request)

	if err == nil {
		return ctx.JSON(http.StatusOK, createdEntity)
	}

	// If there was an error, the error should be returned
	return ctx.JSON(err.(errors.AppError).Code, err.(errors.AppError).Message)
}

func (ctr CrudController) GetAll(ctx echo.Context) error {
	entities, err := ctr.ops.ReadAll()

	if err != nil {
		return ctx.JSON(err.(errors.AppError).Code, err.(errors.AppError).Message)
	}

	return ctx.JSON(http.StatusOK, entities)
}

func (ctr CrudController) Get(ctx echo.Context, id string) error {
	if id == "" {
		e := errors.NewBadRequest("The id should not be empty")
		return ctx.JSON(e.(errors.AppError).Code, e.(errors.AppError).Message)
	}

	entity, err := ctr.ops.Read(id)

	if err != nil {
		return ctx.JSON(err.(errors.AppError).Code, err.(errors.AppError).Message)
	}

	return ctx.JSON(http.StatusOK, entity)
}

func (ctr CrudController) Update(ctx echo.Context) error {
	// Get the attributes from the request body
	var attributes map[string]interface{}
	err := (&echo.DefaultBinder{}).BindBody(ctx, &attributes)

	if err != nil {
		e := errors.NewServerError("Receiving the request data was not successful").(errors.AppError)
		return ctx.JSON(e.Code, e.Message)
	}

	// Check if the body has the correct format
	valid := ctr.validateInput(attributes, ctr.args.Update)
	if !valid {
		msg := fmt.Sprintf("Request body has bad format. It should have the following format: %s", BuildAttributeString(ctr.args.Create))
		e := errors.NewBadRequest(msg)
		return ctx.JSON(e.(errors.AppError).Code, e.(errors.AppError).Message)
	}

	request := ctr.target.BuildRequest(attributes)
	newEntity, err := ctr.ops.Update(request)

	if err != nil {
		return ctx.JSON(err.(errors.AppError).Code, err.(errors.AppError).Message)
	}

	return ctx.JSON(http.StatusOK, newEntity)
}

func (ctr CrudController) Delete(ctx echo.Context, id string) error {
	if id == "" {
		e := errors.NewBadRequest("The id should not be empty")
		return ctx.JSON(e.(errors.AppError).Code, e.(errors.AppError).Message)
	}

	err := ctr.ops.Delete(id)

	if err != nil {
		return ctx.JSON(err.(errors.AppError).Code, err.(errors.AppError).Message)
	}

	return ctx.JSON(http.StatusOK, "")
}
