package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDivide(t *testing.T) {
	t.Run("division witch positive numbers", func(t *testing.T) {
		result, err := Divide(6, 3)
		require.NoError(t, err)
		assert.Equal(t, 2, result)
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Divide(5, 0)
		require.Error(t, err)
	})

	t.Run("division with negative numbers", func(t *testing.T) {
		_, err := Divide(-5, 7)
		require.Error(t, err)
	})

	t.Run("division with indeterminate form", func(t *testing.T) {
		_, err := Divide(0, 0)
		require.Error(t, err)
	})
}
