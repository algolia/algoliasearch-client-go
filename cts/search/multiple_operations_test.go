package search

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestMultipleOperations(t *testing.T) {
	t.Parallel()
	client, _, indexName1 := cts.InitSearchClient1AndIndex(t)
	indexName2 := indexName1 + "_dev"

	var objectIDs []string
	type Book struct{ Title string }
	book := Book{Title: "Harry Potter"}

	{
		operation := search.BatchOperation{
			Action: search.AddObject,
			Body:   book,
		}

		operations := []search.BatchOperationIndexed{
			{IndexName: indexName1, BatchOperation: operation},
			{IndexName: indexName1, BatchOperation: operation},
			{IndexName: indexName2, BatchOperation: operation},
			{IndexName: indexName2, BatchOperation: operation},
		}

		res, err := client.MultipleBatch(operations)
		require.NoError(t, err)
		require.NoError(t, res.Wait())

		objectIDs = res.ObjectIDs
	}

	{
		requests := []search.IndexedGetObject{
			{IndexName: indexName1, ObjectID: objectIDs[0]},
			{IndexName: indexName1, ObjectID: objectIDs[1]},
			{IndexName: indexName2, ObjectID: objectIDs[2]},
			{IndexName: indexName2, ObjectID: objectIDs[3]},
		}
		var books []Book
		err := client.MultipleGetObjects(requests, &books)
		require.NoError(t, err)
		require.Len(t, books, 4)
		for _, b := range books {
			require.Equal(t, book, b)
		}
	}

	{
		queries := []search.IndexedQuery{
			search.NewIndexedQuery(indexName1, opt.Query(""), opt.HitsPerPage(2)),
			search.NewIndexedQuery(indexName2, opt.Query(""), opt.HitsPerPage(2)),
		}

		res, err := client.MultipleQueries(queries, "none")
		require.NoError(t, err)
		require.Len(t, res.Results, 2)
		require.Len(t, res.Results[0].Hits, 2)
		require.Len(t, res.Results[1].Hits, 2)

		res, err = client.MultipleQueries(queries, "stopIfEnoughMatches")
		require.NoError(t, err)
		require.Len(t, res.Results, 2)
		require.Len(t, res.Results[0].Hits, 2)
		require.Len(t, res.Results[1].Hits, 0)
	}
}
