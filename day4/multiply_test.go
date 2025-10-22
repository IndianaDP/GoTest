package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
		wantErr  bool
	}{
		{"multiplication with positive numbers", 2, 3, 6, false},
		{"multiplication with zero", 6, 0, 0, false},
		{"multiplication with indeterminate form", 0, 0, 0, true},
		{"multiplication with negative numbers", -6, 2, 0, true},
		{"multiplication causing overflow", math.MaxInt32, 2, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Multiply(tt.a, tt.b)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
