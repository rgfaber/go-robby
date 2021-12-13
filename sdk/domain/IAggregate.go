package domain

import "go-robby/sdk/schema"

type IAggregate interface {
	Apply(event schema.Event)
	Execute(command schema.Command)
}