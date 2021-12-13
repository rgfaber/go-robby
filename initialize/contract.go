package initialize

import (
	sdk_schema "go-robby/sdk/schema"
)

const GameInitialized 		= "robby.game.initialized"
const InitializeGame      	= "robby.game.initialize"


type Fact struct {
	Topic 		string 						`json:"topic"`
	Meta  		sdk_schema.EventMeta 		`json:"meta"`
	Payload 	[]byte 						`json:"payload"`
}

type Hope struct {
	Topic 		string 						`json:"topic"`
	Meta  		sdk_schema.AggregateMeta 	`json:"meta"`
	Payload 	[]byte 						`json:"payload"`
}

type Feedback struct {
	Meta   		sdk_schema.AggregateMeta 	`json:"meta"`
	Errors 		[]string 					`json:"errors"`
	Payload 	[]byte 						`json:"payload"`
}

func NewFeedback(meta sdk_schema.AggregateMeta, errors []string) *Feedback {
	return &Feedback{Meta: meta, Errors: errors}
}

func (f *Feedback) IsSuccess() (bool,error) {
	return len(f.Errors)==0, nil
}





