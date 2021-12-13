package initialize

import "go-robby/sdk/schema"

type Order struct {
	Id             string        `json:"id"`
	Name           string        `json:"name"`
	NumberOfRobots int           `json:"number_of_robots"`
	Dimensions     schema.Vector `json:"dimensions"`
}

func NewOrder(id string, name string, numberOfRobots int, dimensions schema.Vector) *Order {
	return &Order{
		Id: id,
		Name: name,
		NumberOfRobots: numberOfRobots,
		Dimensions: dimensions}
}






