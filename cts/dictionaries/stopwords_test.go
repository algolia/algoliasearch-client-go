package search

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestStopWords(t *testing.T) {
	t.Parallel()
	client := cts.InitSearchClient2(t)

	entryID := cts.RandSeq(10)
	{
		res, _ := client.SearchDictionaryEntries(search.Stopwords, entryID)
		require.Empty(t, res.Hits)
	}

	stopwordEntry := search.NewStopword(entryID, "en", "down", "enabled")

	{
		res, err := client.SaveDictionaryEntries(search.Stopwords, []search.DictionaryEntry{stopwordEntry})
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	{
		res, err := client.SearchDictionaryEntries(search.Stopwords, entryID)
		require.NoError(t, err)
		require.Equal(t, 1, res.NbHits)
		resEntries, err := res.DictionaryEntires()
		require.NoError(t, err)
		require.Equal(t, stopwordEntry, resEntries[0])
	}

	{
		res, err := client.DeleteDictionaryEntries(search.Stopwords, []string{entryID})
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	{
		res, err := client.SearchDictionaryEntries(search.Stopwords, entryID)
		require.NoError(t, err)
		require.Equal(t, 0, res.NbHits)
	}

	{
		searchRes, err := client.SearchDictionaryEntries(search.Stopwords, "")
		require.NoError(t, err)

		previousEntries, err := searchRes.DictionaryEntires()
		require.NoError(t, err)

		updateRes, err := client.SaveDictionaryEntries(search.Stopwords, []search.DictionaryEntry{stopwordEntry})
		require.NoError(t, err)
		require.NoError(t, updateRes.Wait())

		searchRes, err = client.SearchDictionaryEntries(search.Stopwords, entryID)
		require.NoError(t, err)
		require.Equal(t, 1, searchRes.NbHits)

		updateRes, err = client.ReplaceDictionaryEntries(search.Stopwords, previousEntries)
		require.NoError(t, err)
		require.NoError(t, updateRes.Wait())

		searchRes, err = client.SearchDictionaryEntries(search.Stopwords, entryID)
		require.NoError(t, err)
		require.Empty(t, searchRes.Hits)
	}

	var settings = search.DictionarySettings{
		DisableStandardEntries: opt.DisableStandardEntries(map[string]map[string]bool{"stopwords": {"en": true}}),
	}

	{
		res, err := client.SetDictionarySettings(settings)
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	{
		res, err := client.GetDictionarySettings()
		require.NoError(t, err)
		require.Equal(t, settings, res)
	}

}
