package mathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkError(t *testing.T, tt mathTestCase, operation func(a, b int) (int, error)) {
	t.Helper()
	result, err := operation(tt.a, tt.b)

	if tt.wantErr {
		require.Error(t, err)
	} else {
		require.NoError(t, err)
		assert.Equal(t, tt.expected, result)
	}
}
