package schema

type Weapon int64

const (
	Paper Weapon = 1 << iota
	Rock
	Scissors
)


type Robot struct {
	Id 			string 		`json:"id"`
	Name   string `json:"name"`
	Weapon Weapon `json:"weapon"`
	Health int32  `json:"health"`
}

func NewRobot(id string, name string, weapon Weapon, health int32) *Robot {
	return &Robot{
		Id: id,
		Name: name,
		Weapon: weapon,
		Health: health}
}




