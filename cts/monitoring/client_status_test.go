package monitoring

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestGetCurrentApiStatus(t *testing.T) {
	t.Parallel()
	c := cts.InitMonitoringClient(t)

	s, err := c.GetCurrentApiStatus()
	require.NoError(t, err)
	require.NotNil(t, s)
}

func TestGetCurrentServerStatus(t *testing.T) {
	t.Parallel()
	c := cts.InitMonitoringClient(t)

	s, err := c.GetCurrentServerStatus()
	require.NoError(t, err)
	require.NotNil(t, s)
}
