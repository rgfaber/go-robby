package schema

import (
	"fmt"
	"testing"
)

func TestNewAggregateMeta(t *testing.T) {
	am := NewAggregateMeta("id-00", 0, -1)
	if am==nil {
		t.Fatalf("AggregateMeta could not be created")
	}
	fmt.Printf("Aggregate Meta:\t %+v\n", *am)
}

