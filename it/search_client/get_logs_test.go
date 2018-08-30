package search_client

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
	"github.com/stretchr/testify/require"
)

func TestGetLogs(t *testing.T) {
	t.Parallel()
	c := it.InitSearchClient1(t)

	for i := 0; i < 2; i++ {
		_, err := c.ListIndexes()
		require.NoError(t, err)
	}

	params := algoliasearch.Map{
		"length": 2,
		"offset": 0,
		"type":   "all",
	}

	logs, err := c.GetLogs(params)
	require.NoError(t, err)
	require.Len(t, logs, 2)
}
