package search_client

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"

	"github.com/algolia/algoliasearch-client-go/it"
)

func TestMultipleOperations(t *testing.T) {
	t.Parallel()
	client, _, indexName1 := it.InitSearchClient1AndIndex(t)
	indexName2 := indexName1 + "_dev"

	var objectIDs []string

	{
		operation := algoliasearch.BatchOperation{
			Action: "addObject",
			Body:   algoliasearch.Map{"firstname": "Jimmie"},
		}

		operations := []algoliasearch.BatchOperationIndexed{
			{IndexName: indexName1, BatchOperation: operation},
			{IndexName: indexName1, BatchOperation: operation},
			{IndexName: indexName2, BatchOperation: operation},
			{IndexName: indexName2, BatchOperation: operation},
		}

		res, err := client.MultipleBatch(operations)
		require.NoError(t, err)

		for _, objectID := range res.ObjectIDs {
			objectIDs = append(objectIDs, objectID)
		}

		var wg sync.WaitGroup
		for indexName, taskID := range res.TaskID {
			wg.Add(1)
			go func(wg *sync.WaitGroup, indexName string, taskID int) {
				defer wg.Done()
				it.WaitTasks(t, client.InitIndex(indexName), taskID)
			}(&wg, indexName, taskID)
		}
		wg.Wait()
	}

	{
		requests := []algoliasearch.IndexedGetObject{
			{IndexName: indexName1, ObjectID: objectIDs[0]},
			{IndexName: indexName1, ObjectID: objectIDs[1]},
			{IndexName: indexName2, ObjectID: objectIDs[2]},
			{IndexName: indexName2, ObjectID: objectIDs[3]},
		}
		res, err := client.MultipleGetObjects(requests)
		require.NoError(t, err)
		require.Len(t, res.Results, 4)
	}

	{
		queries := []algoliasearch.IndexedQuery{
			{
				IndexName: indexName1,
				Params:    algoliasearch.Map{"query": "", "hitsPerPage": 2},
			},
			{
				IndexName: indexName2,
				Params:    algoliasearch.Map{"query": "", "hitsPerPage": 2},
			},
		}

		res, err := client.MultipleQueries(queries, "none")
		require.NoError(t, err)
		require.Len(t, res, 2)
		require.Len(t, res[0].Hits, 2)
		require.Len(t, res[1].Hits, 2)

		res, err = client.MultipleQueries(queries, "stopIfEnoughMatches")
		require.NoError(t, err)
		require.Len(t, res, 2)
		require.Len(t, res[0].Hits, 2)
		require.Len(t, res[1].Hits, 0)
	}
}
