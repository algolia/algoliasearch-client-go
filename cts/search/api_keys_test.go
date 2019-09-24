package search

import (
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestAPIKeys(t *testing.T) {
	t.Parallel()
	client := cts.InitSearchClient1(t)

	key := search.Key{
		ACL:                    []string{"search"},
		Description:            "A description",
		MaxHitsPerQuery:        1000,
		MaxQueriesPerIPPerHour: 1000,
		Referers:               []string{"referer"},
		Validity:               10 * time.Minute,
	}

	key.SetQueryParameters(
		opt.TypoToleranceStrict(),
	)

	{
		res, err := client.AddAPIKey(key)
		require.NoError(t, err)
		require.NoError(t, res.Wait())
		key.Value = res.Key
	}

	defer func() {
		_, _ = client.DeleteAPIKey(key.Value)
	}()

	{
		found, err := client.GetAPIKey(key.Value)
		require.NoError(t, err)
		require.True(t, key.Equal(found))
	}

	{
		res, err := client.ListAPIKeys()
		require.NoError(t, err)

		found := false
		for _, k := range res.Keys {
			found = found || k.Value == key.Value
		}
		require.True(t, found)
	}

	key.MaxHitsPerQuery = 42

	{
		res, err := client.UpdateAPIKey(key)
		require.NoError(t, err)
		require.NoError(t, res.Wait())

		found, err := client.GetAPIKey(key.Value)
		require.NoError(t, err)
		require.True(t, key.Equal(found))
	}

	{
		res, err := client.DeleteAPIKey(key.Value)
		require.NoError(t, err)
		require.NoError(t, res.Wait())

		_, err = client.GetAPIKey(key.Value)
		require.Error(t, err)
	}

	{
		res, err := client.RestoreAPIKey(key.Value)
		require.NoError(t, err)
		require.NoError(t, res.Wait())

		found, err := client.GetAPIKey(key.Value)
		require.NoError(t, err)
		require.True(t, key.Equal(found))
	}
}
