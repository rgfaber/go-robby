package schema

import (
	"fmt"
	"testing"
)

func TestNewVector(t *testing.T) {
	v := NewVector(2, 5, 3)
	if v == nil {
		t.Fatalf("No vector was created.")
	}
	if v.X != 2 {
		t.Fatalf("X != 2")
	}
	if v.Y != 5 {
		t.Fatalf("Y != 5")
	}
	if v.Z != 3 {
		t.Fatalf("Z != 3")
	}
	fmt.Printf("Vector:\t\t %+v\n", *v)
}



