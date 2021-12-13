package schema

import (
	"fmt"
	"testing"
)

func TestNewArena(t *testing.T) {
	arena := NewArena(FakeGameId, FakeStatus, *NewFakeVector(), NewFakeRobots(15))
	if arena == nil {
		t.Fatalf("Could not create a new Arena")
	}
	fmt.Println(*arena)
}

