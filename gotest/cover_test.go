package gotest

import "testing"

func add(x,y int) int {
	return x+y
}

func TestCover(t *testing.T) {
	if add(1,2) != 3 {
		t.Fatalf("....")
	}
}
