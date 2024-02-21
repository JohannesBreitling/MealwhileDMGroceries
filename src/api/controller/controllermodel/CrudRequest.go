package controllermodel

import "mealwhile/logic/model"

type CrudRequest interface {
	Build(id bool) model.CrudEntity
}
