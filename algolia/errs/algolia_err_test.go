package errs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsAlgoliaErr(t *testing.T) {
	actual, ok := IsAlgoliaErr(nil)
	require.Nil(t, actual)
	require.False(t, ok)

	_, ok = IsAlgoliaErr(fmt.Errorf("random error"))
	require.False(t, ok)

	_, ok = IsAlgoliaErr(&AlgoliaErr{})
	require.True(t, ok)

	_, ok = IsAlgoliaErr(AlgoliaErr{})
	require.True(t, ok)
}
