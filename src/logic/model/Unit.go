package model

type Unit struct {
	Id           string
	Name         string
	Abbreviation string
}

func (u Unit) GetId() string {
	return u.Id
}
