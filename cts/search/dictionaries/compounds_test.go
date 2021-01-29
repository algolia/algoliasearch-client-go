package search

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestCompounds(t *testing.T) {
	t.Parallel()
	client := cts.InitSearchClient2(t)

	entryID := cts.GenerateRandomString(10)
	{
		res, _ := client.SearchDictionaryEntries(search.Compounds, entryID)
		require.Empty(t, res.Hits)
	}

	compoundEntry := search.NewCompound(entryID, "de", "kopfschmerztablette", []string{"kopf", "schmerz", "tablette"})

	{
		res, err := client.SaveDictionaryEntries(search.Compounds, []search.DictionaryEntry{compoundEntry})
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	{
		res, err := client.SearchDictionaryEntries(search.Compounds, entryID)
		require.NoError(t, err)
		require.Equal(t, 1, res.NbHits)
		resEntries, err := res.DictionaryEntries()
		require.NoError(t, err)
		require.Equal(t, compoundEntry, resEntries[0])
	}

	{
		res, err := client.DeleteDictionaryEntries(search.Compounds, []string{entryID})
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	{
		res, err := client.SearchDictionaryEntries(search.Compounds, entryID)
		require.NoError(t, err)
		require.Equal(t, 0, res.NbHits)
	}

}
