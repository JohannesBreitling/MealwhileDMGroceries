package data

type GroceryRepositoryInterface interface {
	CrudRepositoryInterface
	FlagReferenced(flagId string) (bool, error)
}
