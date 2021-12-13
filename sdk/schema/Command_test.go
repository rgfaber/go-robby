package schema

import (
	"encoding/json"
	"fmt"
	"go-robby/app/schema"
	"testing"
	"time"
)

func TestNewCommand(t *testing.T) {
	fakePayload := schema.NewFakePayload(schema.FakeGameId, schema.FakeGameName, 42, time.Now())
	fakeAggregateMeta := schema.NewFakeAggregateMeta()
	fakePayload = schema.NewFakePayload(schema.FakeGameId, schema.FakeGameName, 42, time.Now().UTC())
	fakeBytes, err := json.Marshal(fakePayload)
	if err!=nil {
		t.Fatalf("Error: \t %+v\n", err)
	}
	cmd := NewCommand(schema.FakeTopic, *fakeAggregateMeta, fakeBytes )
	if cmd==nil {
		t.Fatalf("Error creating Command")
	}
	fmt.Printf("Command:\t %+v\n", *cmd)
}