package schema

import (
	"go-robby/sdk/schema"
	"reflect"
	"time"
)

var Nouns=[5]string {"Angel","Bandit","Rascal","Misfit","Robber"}
var Adjectives = [5]string {""}




const (
	FakeGameId 			= "fake-game00"
	FakeGameName 		= "fake-game-name"
	FakeNumberOfRobots 	= 4
	FakeX 				= 15
	FakeY 				= 20
	FakeZ 				= 42
	FakeEventId			= "fake-event-00"
	FakeEventType		= "fake-event-type"
	FakeTopic 			= "eventmeta.tested"
)

type FakePayload struct {
	Id 			string 		`json:"id"`
	Name 		string 		`json:"name"`
	Weight 		int64 		`json:"weight"`
	SomeMoment 	time.Time 	`json:"some_moment"`
}

func NewFakePayload(id string, name string, weight int64, someMoment time.Time) *FakePayload {
	return &FakePayload{
		Id: id,
		Name: name,
		Weight: weight,
		SomeMoment: someMoment}
}

const FakeStatus Status = Pending


func NewFakeAggregateMeta() *schema.AggregateMeta {
	return schema.NewAggregateMeta(FakeGameId, int64(FakeStatus), -1)
}

func NewFakeVector() *schema.Vector {
	return schema.NewVector(FakeX, FakeY, FakeZ)
}

func NewFakeEventMeta() *schema.EventMeta {
	ev := schema.Event{}
	return schema.NewEventMeta(FakeEventId, time.Now(), reflect.TypeOf(ev).String(), FakeTopic)
}

func NewFakeRobots(count int) []Robot {
  return make([]Robot, count)
}


