package schema

import (
	"fmt"
	"testing"
)

func TestNewRobot(t *testing.T) {
	r := NewRobot("robot-000", "Robby", Scissors, 42 )
	if r==nil {
		t.Fatalf("Robot could not be created")
	}
	fmt.Printf("Robot:\t\t %+v\n", *r)
}
