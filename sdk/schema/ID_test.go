package schema

import (
	"fmt"
	"go-robby/app/schema"
	"testing"
	"time"
)


func TestNewID(t *testing.T) {
	id := NewID(schema.FakeGameId)
	if id==nil {
		t.Fatalf("Could not create Id")
	}
	fmt.Printf("Test Id:\t %+v\n", *id)
}


func TestUtc(t *testing.T) {
	s := time.UTC.String()
	fmt.Println(s)
}