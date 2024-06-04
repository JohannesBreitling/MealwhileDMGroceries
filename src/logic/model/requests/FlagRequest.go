package requests

type FlagRequest struct {
	Id          string
	Name        string
	Description string
}

func (r FlagRequest) GetId() string {
	return r.Id
}
