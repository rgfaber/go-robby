package schema

import (
	"fmt"
	"go-robby/app/schema"
	"reflect"
	"testing"
	"time"
)

const (

)

func TestNewEventMeta(t *testing.T) {
	var e = EventMeta{}
	evenType := fmt.Sprintf("%+v", reflect.ValueOf(e).Kind())
	em := NewEventMeta(schema.FakeEventId, time.Now(), evenType, schema.FakeTopic)
	if &em==nil {
		t.Fatalf("Event Metadata could not be created")
	}
	fmt.Printf("Event Meta:\t %+v\n", *em)

}
