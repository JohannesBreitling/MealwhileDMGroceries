package model

type CrudEntity interface {
	GetId() string
	SetId(id string)

	Empty() CrudEntity
	Attributes() map[string]string
	String() string
	EntityName() string
	FromArguments(map[string]string) CrudEntity
	FromInterface(map[string]interface{}) CrudEntity
}
