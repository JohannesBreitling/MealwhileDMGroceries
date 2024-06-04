package requests

type GroceryRequest struct {
	Id    string
	Name  string
	Flags []string
}

func (r GroceryRequest) GetId() string {
	return r.Id
}
