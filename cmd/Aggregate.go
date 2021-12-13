package robby

import (
	"encoding/json"
	schema2 "go-robby/app/schema"
	"go-robby/initialize"
	"go-robby/sdk/schema"
)


type Aggregate struct {
	Model schema2.Arena `json:"schema"`
}

func NewAggregate(model schema2.Arena) *Aggregate {
	return &Aggregate{Model: model}
}


func (agg *Aggregate) Execute(command *schema.Command) {

}


func (agg *Aggregate) Apply(event *schema.Event) {
	if event==nil {panic(event)}
	switch event.Meta.Topic {
	case initialize.GameInitialized:
		err := json.Unmarshal(event.Payload, agg.Model.Dimensions)
		if err!=nil {
			panic(agg.Model.Dimensions)
		}
		agg.Model.Status = schema2.Initialized
	}
	
}

func (agg *Aggregate) FromEvents(events []schema.Event)  *Aggregate {
	for i := range events {
		agg.Apply(&events[i])
	}
	return agg
}

