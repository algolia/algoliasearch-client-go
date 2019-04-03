package search_index

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/algolia/algoliasearch-client-go/cts"
	"github.com/stretchr/testify/require"
)

func TestBatching(t *testing.T) {
	t.Parallel()
	_, index, _ := cts.InitSearchClient1AndIndex(t)

	{
		res, err := index.SaveObjects([]map[string]string{
			{"objectID": "one", "key": "value"},
			{"objectID": "two", "key": "value"},
			{"objectID": "three", "key": "value"},
			{"objectID": "four", "key": "value"},
			{"objectID": "five", "key": "value"},
		})
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	{
		res, err := index.Batch([]search.BatchOperation{
			{Action: search.AddObject, Body: map[string]string{"objectID": "zero", "key": "value"}},
			{Action: search.UpdateObject, Body: map[string]string{"objectID": "one", "k": "v"}},
			{Action: search.PartialUpdateObject, Body: map[string]string{"objectID": "two", "k": "v"}},
			{Action: search.PartialUpdateObject, Body: map[string]string{"objectID": "two_bis", "key": "value"}},
			{Action: search.PartialUpdateObjectNoCreate, Body: map[string]string{"objectID": "three", "k": "v"}},
			{Action: search.DeleteObject, Body: map[string]string{"objectID": "four"}},
		})
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	{
		var found []map[string]string
		expected := []map[string]string{
			{"objectID": "zero", "key": "value"},
			{"objectID": "one", "k": "v"},
			{"objectID": "two", "key": "value", "k": "v"},
			{"objectID": "two_bis", "key": "value"},
			{"objectID": "three", "key": "value", "k": "v"},
			{"objectID": "five", "key": "value"},
		}

		it, err := index.BrowseObjects()
		require.NoError(t, err)

		for {
			var hit map[string]string
			_, err := it.Next(&hit)
			if err != nil {
				require.Equal(t, search.NoMoreHitsErr, err)
				break
			}
			found = append(found, hit)
		}

		require.ElementsMatch(t, expected, found)
	}
}
