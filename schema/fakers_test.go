package schema

import (
	"fmt"
	"testing"
)


func TestNewFakeAggregateMeta(t *testing.T) {
	em := NewFakeAggregateMeta()
	if em==nil {
		t.Fatalf("Fake Aggregate Meta could not be created.")
	}
	fmt.Printf("Fake Agg.Meta:\t %+v\n", *em)
}


func TestNewFakeVector(t *testing.T) {
	fv := NewFakeVector()
	if fv==nil {
		t.Fatalf("Could not create fake Vector")
	}
	fmt.Printf("Fake Vector:\t %+v\n", *fv)
}
