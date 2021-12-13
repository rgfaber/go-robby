package schema

import "go-robby/sdk/schema"


type Status int64

const (
	Pending Status = 0 << iota
	Initialized
)



type Arena struct {
	Id         string        `json:"id"`
	Status     Status        `json:"status"`
	Dimensions schema.Vector `json:"dimensions"`
	Robots     []Robot       `json:"robots"`
}

func NewArena(
	id 			string,
	status Status,
	dimensions 	schema.Vector,
	robots 		[]Robot,
) *Arena {
	return &Arena{
		Id: id,
		Status: status,
		Dimensions: dimensions,
		Robots: robots,
	}
}
