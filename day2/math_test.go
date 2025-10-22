package main

import (
	"testing"
)

func TestDivide(t *testing.T) {

	result, err := Divide(6, 3)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != 2 {
		t.Errorf("expected result to be 2, got %d", result)
	}

	_, err = Divide(5, 0)

	if err == nil {
		t.Fatalf("expected an error fo division by zero, go nil")
	}

	_, err = Divide(0, 0)

	if err == nil {
		t.Fatalf("expected an error for indeterminate form 0/0, got nil")
	}

	_, err = Divide(-1, 2)

	if err == nil {
		t.Fatalf("expected an error for negative numbers, got nil")
	}
}
