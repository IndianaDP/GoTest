package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDivide(t *testing.T) {

	result, err := Divide(10, 2)

	require.NoError(t, err, "expected no error, got %v", err)
	assert.Equal(t, 5, result, "expected result to be 5, got %d", result)

	_, err = Divide(5, 0)

	require.Error(t, err, "expected an error for division by zero, got nil")

	_, err = Divide(0, 0)

	require.Error(t, err, "expected an error for indeterminate form 0/0, got nil")

	_, err = Divide(-1, 2)

	require.Error(t, err, "expected an error for negative numbers, got nil")
}
