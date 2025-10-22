package main

import "testing"

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected: 2, got: %d", result)
	}

	result = Subtract(2, 5)
	if result != -3 {
		t.Errorf("Expected: -3, got: %d", result)
	}
}
