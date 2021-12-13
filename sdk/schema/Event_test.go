package schema

import (
	"encoding/json"
	"fmt"
	"go-robby/app/schema"
	"testing"
)


type testLoad struct {
	Id 		string 		`json:"id"`
	Name	string 		`json:"name"`
	Weight 	int 		`json:"weight"`
}

func newTestLoad(id string, name string, weight int) *testLoad {
	return &testLoad{Id: id, Name: name, Weight: weight}
}


func TestNewEvent(t *testing.T) {
	tl := newTestLoad(schema.FakeGameId, schema.FakeGameName, 42)
	tlb, err := json.Marshal(tl)
	if err!=nil {
		t.Fatalf("json.Marshal error %+v\n", err)
	}
	evm := schema.NewFakeEventMeta()
	ev := NewEvent(tlb, evm)
	if ev == nil {
		t.Fatalf("Event could not be created")
	}
	fmt.Printf("Event:\t %+v", *ev)
}