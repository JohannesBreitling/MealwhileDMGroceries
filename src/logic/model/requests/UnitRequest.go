package requests

type UnitRequest struct {
	Id           string
	Name         string
	Abbreviation string
}

func (r UnitRequest) GetId() string {
	return r.Id
}
