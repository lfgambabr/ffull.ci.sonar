package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(5, 6)

	if result != 11 {
		t.Error("The result must be 11")
	}
}
