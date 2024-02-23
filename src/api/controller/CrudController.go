package controller

import (
	"fmt"
	"mealwhile/logic/model"
	"mealwhile/logic/operations/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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
	args ExpectedCrudArguments
	ops  interfaces.CrudServiceInterface
}

func NewCrudController(ops interfaces.CrudServiceInterface, args ExpectedCrudArguments) CrudController {
	return CrudController{ops: ops, args: args}
}

func (CrudController) validateInput(input map[string]string, expected []string) bool {
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

func (ctr CrudController) Create(ctx echo.Context, entity model.CrudEntity) error {
	var msg string

	// Get the attributes from the request body
	var attributes map[string]string
	err := ctx.Bind(&attributes)

	// Check if the body has the correct format
	valid := ctr.validateInput(attributes, ctr.args.Create)
	msg = fmt.Sprintf("request body has bad format. It should have the following format: %s", BuildAttributeString(ctr.args.Create))
	if err != nil || !valid {
		log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}

	entityBuilt := entity.FromArguments(attributes)

	// Create the entity
	createdEntity, err := ctr.ops.Create(entityBuilt)

	// If there was an error, the error should be returned
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, createdEntity)
}

func (ctr CrudController) GetAll(ctx echo.Context, target model.CrudEntity) error {
	entities, err := ctr.ops.ReadAll(target)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, entities)
}

func (ctr CrudController) Get(ctx echo.Context, target model.CrudEntity, id string) error {
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, "the id should not be empty")
	}

	entity, err := ctr.ops.Read(target, id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, entity)
}

func (ctr CrudController) Update(ctx echo.Context, entity model.CrudEntity) error {
	var msg string

	// Get the attributes from the request body
	var attributes map[string]string
	err := (&echo.DefaultBinder{}).BindBody(ctx, &attributes)

	// Check if the body has the correct format
	valid := ctr.validateInput(attributes, ctr.args.Update)
	msg = fmt.Sprintf("request body has bad format. It should have the following format: %s", BuildAttributeString(ctr.args.Update))
	if err != nil || !valid {
		log.Error(msg)
		return ctx.JSON(http.StatusBadRequest, msg)
	}

	entityBuilt := entity.FromArguments(attributes)

	newEntity, err := ctr.ops.Update(entityBuilt)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, newEntity)
}

func (ctr CrudController) Delete(ctx echo.Context, target model.CrudEntity, id string) error {
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, "the id should not be empty")
	}

	err := ctr.ops.Delete(target, id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "")
}
