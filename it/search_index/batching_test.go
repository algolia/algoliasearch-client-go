package search_index

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
)

func TestBatching(t *testing.T) {
	t.Parallel()
	_, index, _ := it.InitSearchClient1AndIndex(t)

	{
		res, err := index.AddObjects([]algoliasearch.Object{
			{"objectID": "one", "key": "value"},
			{"objectID": "two", "key": "value"},
			{"objectID": "three", "key": "value"},
			{"objectID": "four", "key": "value"},
			{"objectID": "five", "key": "value"},
		})
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)
	}

	{
		res, err := index.Batch([]algoliasearch.BatchOperation{
			{Action: "addObject", Body: algoliasearch.Object{"objectID": "zero", "key": "value"}},
			{Action: "updateObject", Body: algoliasearch.Object{"objectID": "one", "k": "v"}},
			{Action: "partialUpdateObject", Body: algoliasearch.Object{"objectID": "two", "k": "v"}},
			{Action: "partialUpdateObject", Body: algoliasearch.Object{"objectID": "two_bis", "key": "value"}},
			{Action: "partialUpdateObjectNoCreate", Body: algoliasearch.Object{"objectID": "three", "k": "v"}},
			{Action: "deleteObject", Body: algoliasearch.Object{"objectID": "four"}},
		})
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)
	}

	{
		var found []algoliasearch.Map
		expected := []algoliasearch.Map{
			{"objectID": "zero", "key": "value"},
			{"objectID": "one", "k": "v"},
			{"objectID": "two", "key": "value", "k": "v"},
			{"objectID": "two_bis", "key": "value"},
			{"objectID": "three", "key": "value", "k": "v"},
			{"objectID": "five", "key": "value"},
		}

		it, err := index.BrowseAll(nil)
		require.NoError(t, err)

		for {
			hit, err := it.Next()
			if err != nil {
				require.Equal(t, algoliasearch.NoMoreHitsErr, err)
				break
			}
			found = append(found, hit)
		}

		require.ElementsMatch(t, expected, found)
	}
}
