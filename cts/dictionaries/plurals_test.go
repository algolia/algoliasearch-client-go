package search

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestPlurals(t *testing.T) {
	t.Parallel()
	client := cts.InitSearchClient2(t)

	entryID := cts.RandSeq(10)
	{
		res, _ := client.SearchDictionaryEntries(search.Plurals, entryID)
		require.Empty(t, res.Hits)
	}

	pluralEntry := search.NewPlural(entryID, "fr", []string{"cheval", "chevaux"})

	{
		res, err := client.SaveDictionaryEntries(search.Plurals, []search.DictionaryEntry{pluralEntry})
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	// {
	// 	res, err := client.SearchDictionaryEntries(search.Plurals, entryID)
	// 	require.NoError(t, err)
	// 	require.Equal(t, 1, res.NbHits)
	// 	resEntries, err := res.DictionaryEntires()
	// 	require.NoError(t, err)
	// 	require.Equal(t, pluralEntry, resEntries[0])
	// }

	// {
	// 	res, err := client.DeleteDictionaryEntries(search.Plurals, []string{entryID})
	// 	require.NoError(t, err)
	// 	require.NoError(t, res.Wait())
	// }

	// {
	// 	res, err := client.SearchDictionaryEntries(search.Plurals, entryID)
	// 	require.NoError(t, err)
	// 	require.Equal(t, 0, res.NbHits)
	// }
}
