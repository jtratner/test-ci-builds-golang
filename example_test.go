package main

import "testing"

func TestNoOp(t *testing.T) {
	if false {
		t.Errorf("Got false")
	}
}
