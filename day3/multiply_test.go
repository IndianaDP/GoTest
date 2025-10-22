package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiply(t *testing.T) {
	result, err := Multiply(2, 3)

	require.NoError(t, err, "expected no error, got %v", err)
	assert.Equal(t, 6, result, "expected result to be 6, got %d", result)

	result, err = Multiply(6, 0)

	require.NoError(t, err, "expected no error, got %v", err)
	assert.Equal(t, 0, result, "expected result to be 0< got %d", result)

	_, err = Multiply(0, 0)
	require.Error(t, err, "expected an error for indeterminate form 0*0, got nil")

	_, err = Multiply(-6, 2)
	require.Error(t, err, "expected an error for negative numbers, got nil")
}
