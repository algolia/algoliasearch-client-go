package search_index

import (
	"fmt"
	"sync"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
	"github.com/stretchr/testify/require"
)

func TestIndexing(t *testing.T) {
	t.Parallel()
	_, index, _ := it.InitSearchClient1AndIndex(t)

	var taskIDs []int
	var objectIDs []string

	{
		res, err := index.AddObject(algoliasearch.Object{"objectID": "one", "attribute": "value"})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
		objectIDs = append(objectIDs, res.ObjectID)

		res, err = index.AddObject(algoliasearch.Object{"attribute": "value"})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
		objectIDs = append(objectIDs, res.ObjectID)
	}

	{
		res, err := index.AddObjects([]algoliasearch.Object{
			algoliasearch.Object{"objectID": "two", "attribute": "value"},
			algoliasearch.Object{"objectID": "three", "attribute": "value"},
		})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
		objectIDs = append(objectIDs, res.ObjectIDs...)

		res, err = index.AddObjects([]algoliasearch.Object{
			algoliasearch.Object{"attribute": "value"},
			algoliasearch.Object{"attribute": "value"},
		})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
		objectIDs = append(objectIDs, res.ObjectIDs...)
	}

	{
		for i := 0; i < 10; i++ {
			var operations []algoliasearch.BatchOperation
			for j := 0; j < 100; j++ {
				operations = append(operations, algoliasearch.BatchOperation{
					Action: "addObject",
					Body:   algoliasearch.Object{"objectID": fmt.Sprintf("%d", i*100+j), "attribute": "value"},
				})
			}
			res, err := index.Batch(operations)
			require.NoError(t, err)
			taskIDs = append(taskIDs, res.TaskID)
			objectIDs = append(objectIDs, res.ObjectIDs...)
		}
	}

	it.WaitTasks(t, index, taskIDs...)
	taskIDs = []int{}

	var expected []algoliasearch.Object
	for _, objectID := range objectIDs {
		expected = append(expected, algoliasearch.Object{"objectID": objectID, "attribute": "value"})
	}

	{
		var wg sync.WaitGroup

		for _, object := range expected[:7] {
			wg.Add(1)
			go getObjectAndCompareWith(t, &wg, index, object)
		}

		wg.Add(1)
		go getObjectsAndCompareWith(t, &wg, index, objectIDs[7:], expected[7:])

		wg.Wait()
	}

	{
		it, err := index.BrowseAll(nil)
		require.NoError(t, err)

		var found []algoliasearch.Object
		i := 0
		for {
			res, err := it.Next()
			i++
			if err == algoliasearch.NoMoreHitsErr {
				break
			}
			require.NoError(t, err)
			found = append(found, algoliasearch.Object(res))
		}
		require.ElementsMatch(t, expected, found)
	}

	{
		var found []algoliasearch.Object
		var cursor string

		for {
			res, err := index.Browse(nil, cursor)
			require.NoError(t, err)
			for _, object := range res.Hits {
				found = append(found, algoliasearch.Object(object))
			}
			if len(res.Hits) < res.HitsPerPage || res.Cursor == "" {
				break
			}
			cursor = res.Cursor
		}
		require.ElementsMatch(t, expected, found)
	}

	{
		res, err := index.UpdateObject(algoliasearch.Object{"objectID": "one", "new_attribute": "new_value"})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index.UpdateObjects([]algoliasearch.Object{
			algoliasearch.Object{"objectID": "two", "new_attribute": "new_value"},
			algoliasearch.Object{"objectID": "three", "new_attribute": "new_value"},
		})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index.PartialUpdateObject(algoliasearch.Object{"objectID": "one", "extra_attribute": "extra_value"})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index.PartialUpdateObjects([]algoliasearch.Object{
			algoliasearch.Object{"objectID": "two", "extra_attribute": "extra_value"},
			algoliasearch.Object{"objectID": "three", "extra_attribute": "extra_value"},
		})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	it.WaitTasks(t, index, taskIDs...)
	taskIDs = []int{}

	{
		var wg sync.WaitGroup
		wg.Add(3)
		go getObjectAndCompareWith(t, &wg, index, algoliasearch.Object{"objectID": "one", "new_attribute": "new_value", "extra_attribute": "extra_value"})
		go getObjectAndCompareWith(t, &wg, index, algoliasearch.Object{"objectID": "two", "new_attribute": "new_value", "extra_attribute": "extra_value"})
		go getObjectAndCompareWith(t, &wg, index, algoliasearch.Object{"objectID": "three", "new_attribute": "new_value", "extra_attribute": "extra_value"})
		wg.Wait()
	}

	{
		for _, objectID := range objectIDs[:7] {
			res, err := index.DeleteObject(objectID)
			require.NoError(t, err)
			taskIDs = append(taskIDs, res.TaskID)
		}
	}

	{
		res, err := index.DeleteObjects(objectIDs[7:])
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	it.WaitTasks(t, index, taskIDs...)
	taskIDs = []int{}

	{
		_, err := index.BrowseAll(nil)
		require.Equal(t, algoliasearch.NoMoreHitsErr, err)
	}
}

func getObjectAndCompareWith(t *testing.T, wg *sync.WaitGroup, index algoliasearch.Index, expected algoliasearch.Object) {
	defer wg.Done()
	objectID, err := expected.ObjectID()
	require.NoError(t, err)
	object, err := index.GetObject(objectID, nil)
	require.NoError(t, err)
	require.Equal(t, expected, object)
}

func getObjectsAndCompareWith(t *testing.T, wg *sync.WaitGroup, index algoliasearch.Index, objectIDs []string, expected []algoliasearch.Object) {
	defer wg.Done()
	objects, err := index.GetObjects(objectIDs)
	require.NoError(t, err)
	require.ElementsMatch(t, expected, objects)
}
