package initialize

import (
	"fmt"
	"go-robby/app/schema"
	"testing"
)

func TestNewFeedback(t *testing.T) {
	fb := NewFeedback(*schema.NewFakeAggregateMeta(), nil)
	if fb==nil {
		t.Fatalf("Could not create Feedback")
	}
	fmt.Printf("Feedback: \t %+v\n", *fb)
}

