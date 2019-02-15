package search_client

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/it"
)

func TestAPIKeys(t *testing.T) {
	t.Parallel()
	client := it.InitSearchClient1(t)

	key := algoliasearch.Key{
		ACL:                    []string{"search"},
		CreatedAt:              0,
		Description:            "A description",
		Indexes:                []string{"index"},
		MaxHitsPerQuery:        1000,
		MaxQueriesPerIPPerHour: 1000,
		QueryParameters:        "typoTolerance=strict",
		Referers:               []string{"referer"},
		Validity:               600,
		Value:                  "",
	}
	acl := key.ACL
	params := algoliasearch.Map{
		"description":            key.Description,
		"indexes":                key.Indexes,
		"maxHitsPerQuery":        key.MaxHitsPerQuery,
		"maxQueriesPerIPPerHour": key.MaxQueriesPerIPPerHour,
		"queryParameters":        key.QueryParameters,
		"referers":               key.Referers,
		"validity":               key.Validity,
	}

	{
		res, err := client.AddAPIKey(acl, params)
		require.NoError(t, err)
		key.Value = res.Key
	}

	defer client.DeleteAPIKey(key.Value)

	{
		it.Retry(func() bool {
			_, err := client.GetAPIKey(key.Value)
			return err == nil
		})

		found, err := client.GetAPIKey(key.Value)
		require.NoError(t, err)

		key.CreatedAt = found.CreatedAt
		key.Validity = found.Validity
		require.Equal(t, key, found)
	}

	{
		keys, err := client.ListAPIKeys()
		require.NoError(t, err)

		found := false
		for _, k := range keys {
			found = k.Value == key.Value
			if found {
				break
			}
		}
		require.True(t, found)
	}

	{
		key.MaxHitsPerQuery = 42
		params = algoliasearch.Map{"maxHitsPerQuery": key.MaxHitsPerQuery}
		_, err := client.UpdateAPIKey(key.Value, params)
		require.NoError(t, err)

		it.Retry(func() bool {
			found, err := client.GetAPIKey(key.Value)
			require.NoError(t, err)
			return found.MaxHitsPerQuery == key.MaxHitsPerQuery
		})
	}

	{
		_, err := client.DeleteAPIKey(key.Value)
		require.NoError(t, err)

		it.Retry(func() bool {
			_, err := client.GetAPIKey(key.Value)
			return err != nil
		})
	}

	{
		_, err := client.RestoreAPIKey(key.Value)
		require.NoError(t, err)

		it.Retry(func() bool {
			_, err := client.GetAPIKey(key.Value)
			return err == nil
		})
	}
}
