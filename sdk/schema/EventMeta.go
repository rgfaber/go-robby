package schema

import (
	"time"
)

type EventMeta struct {
	Id string `json:"id"`
	TimeStamp time.Time `json:"time_stamp"`
	Type string `json:"type"`
	Topic string
}

func NewEventMeta(id string, timeStamp time.Time,typeName string, topic string) *EventMeta {
return &EventMeta{
	Id: id,
	TimeStamp: timeStamp,
	Type: typeName,
	Topic: topic}
}


