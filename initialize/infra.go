package initialize

import (
	"go-robby/sdk/schema"
)


type Emitter struct {
	channel chan schema.Event
}

func NewEmitter(channel chan schema.Event) *Emitter {
	return &Emitter{channel: channel}
}




func (e *Emitter) Emit(event schema.Event) {
	e.channel <- event
}