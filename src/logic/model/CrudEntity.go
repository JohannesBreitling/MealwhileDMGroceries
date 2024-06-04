package model

import "mealwhile/logic/model/requests"

type CrudEntity interface {
	GetId() string
	SetId(id string)

	Empty() CrudEntity
	Attributes() map[string]interface{}
	String() string
	EntityName() string
	FromInterface(map[string]interface{}) CrudEntity
	BuildRequest(map[string]interface{}) requests.CrudRequest
}
