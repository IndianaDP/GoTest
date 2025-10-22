package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDivideSub(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
		wantErr  bool
	}{
		{"division with positive numbers", 6, 3, 2, false},
		{"division by zero", 5, 0, 0, true},
		{"division with negative numbers", -5, 7, 0, true},
		{"division with indeterminate form", 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
