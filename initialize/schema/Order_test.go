package schema

import (
	"fmt"
	"go-robby/sdk/schema"
	"testing"
)


func createTestOrder() *Order {
	v := schema.NewVector(5,4,7)
	g := NewOrder("game-001", "my game", 3, *v )
	return g
}



func TestNewOrder(t *testing.T) {
	g := createTestOrder()
	if &g == nil{
		t.Fatalf(`Could not create a game`)
	}
}

func TestSerialization(t *testing.T) {
	g := createTestOrder()
	fmt.Printf("Order:\t\t %+v\n" , *g)
}
