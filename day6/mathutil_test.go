package mathutil

import (
	"math"
	"testing"
)

type mathTestCase struct {
	name     string
	a, b     int
	expected int
	wantErr  bool
}

func TestMultiply(t *testing.T) {
	tests := []mathTestCase{
		{"multiplication with positive numbers", 2, 3, 6, false},
		{"multiplication with zero", 6, 0, 0, false},
		{"multiplication with indeterminate form", 0, 0, 0, true},
		{"multiplication with negative numbers", -6, 2, 0, true},
		{"multiplication causing overflow", math.MaxInt32, 2, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkError(t, tt, Multiply)
		})
	}
}

func TestDivideSub(t *testing.T) {
	tests := []mathTestCase{
		{"division with positive numbers", 6, 3, 2, false},
		{"division by zero", 5, 0, 0, true},
		{"division with negative numbers", -5, 7, 0, true},
		{"division with indeterminate form", 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkError(t, tt, Divide)
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []mathTestCase{
		{"addition with positive numbers", 2, 3, 5, false},
		{"addition with indeterminate form", 0, 0, 0, true},
		{"addition with negative numbers", -2, 3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkError(t, tt, Add)
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []mathTestCase{
		{"subtraction with positive numbers", 5, 3, 2, false},
		{"subtraction with indeterminate form", 0, 0, 0, true},
		{"subtraction with negative numbers", -5, 3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkError(t, tt, Subtract)
		})
	}
}
