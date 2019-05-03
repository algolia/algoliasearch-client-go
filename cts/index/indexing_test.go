package index

import (
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/algolia/algoliasearch-client-go/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/cts"
	"github.com/stretchr/testify/require"
)

func TestIndexing(t *testing.T) {
	t.Parallel()
	_, index, _ := cts.InitSearchClient1AndIndex(t)

	g := wait.NewGroup()
	var objectIDs []string

	{
		res, err := index.SaveObject(map[string]string{"objectID": "one", "attribute": "value"})
		require.NoError(t, err)
		g.Collect(res)
		objectIDs = append(objectIDs, res.ObjectID)

		res, err = index.SaveObject(map[string]string{"attribute": "value"})
		require.NoError(t, err)
		g.Collect(res)
		objectIDs = append(objectIDs, res.ObjectID)
	}

	{
		res, err := index.SaveObjects(nil, opt.AutoGenerateObjectIDIfNotExist(true))
		require.NoError(t, err)
		g.Collect(res)

		res, err = index.SaveObjects([]map[string]interface{}{}, opt.AutoGenerateObjectIDIfNotExist(true))
		require.NoError(t, err)
		g.Collect(res)
	}

	{
		res, err := index.SaveObjects([]map[string]string{
			{"objectID": "two", "attribute": "value"},
			{"objectID": "three", "attribute": "value"},
		})
		require.NoError(t, err)
		g.Collect(res)
		objectIDs = append(objectIDs, res.ObjectIDs()...)

		res, err = index.SaveObjects([]map[string]string{
			{"attribute": "value"},
			{"attribute": "value"},
		}, opt.AutoGenerateObjectIDIfNotExist(true))
		require.NoError(t, err)
		g.Collect(res)
		objectIDs = append(objectIDs, res.ObjectIDs()...)
	}

	{
		for i := 0; i < 10; i++ {
			var operations []search.BatchOperation
			for j := 0; j < 100; j++ {
				operations = append(operations, search.BatchOperation{
					Action: "addObject",
					Body:   map[string]string{"objectID": fmt.Sprintf("%d", i*100+j), "attribute": "value"},
				})
			}
			res, err := index.Batch(operations)
			require.NoError(t, err)
			g.Collect(res)
			objectIDs = append(objectIDs, res.ObjectIDs...)
		}
	}

	require.NoError(t, g.Wait())

	var expected []map[string]string
	for _, objectID := range objectIDs {
		expected = append(expected, map[string]string{"objectID": objectID, "attribute": "value"})
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
		it, err := index.BrowseObjects()
		require.NoError(t, err)

		var found []map[string]string
		i := 0
		for {
			var object map[string]string
			_, err := it.Next(&object)
			i++
			if err == io.EOF {
				break
			}
			require.NoError(t, err)
			found = append(found, object)
		}
		require.ElementsMatch(t, expected, found)
	}

	{
		res, err := index.SaveObject(map[string]string{"objectID": "one", "new_attribute": "new_value"})
		require.NoError(t, err)
		g.Collect(res)
	}

	{
		res, err := index.SaveObjects([]map[string]string{
			{"objectID": "two", "new_attribute": "new_value"},
			{"objectID": "three", "new_attribute": "new_value"},
		})
		require.NoError(t, err)
		g.Collect(res)
	}

	{
		res, err := index.PartialUpdateObject(map[string]string{"objectID": "one", "extra_attribute": "extra_value"})
		require.NoError(t, err)
		g.Collect(res)
	}

	{
		res, err := index.PartialUpdateObjects([]map[string]string{
			{"objectID": "two", "extra_attribute": "extra_value"},
			{"objectID": "three", "extra_attribute": "extra_value"},
		})
		require.NoError(t, err)
		g.Collect(res)
	}

	require.NoError(t, g.Wait())

	{
		var wg sync.WaitGroup
		wg.Add(3)
		go getObjectAndCompareWith(t, &wg, index, map[string]string{"objectID": "one", "new_attribute": "new_value", "extra_attribute": "extra_value"})
		go getObjectAndCompareWith(t, &wg, index, map[string]string{"objectID": "two", "new_attribute": "new_value", "extra_attribute": "extra_value"})
		go getObjectAndCompareWith(t, &wg, index, map[string]string{"objectID": "three", "new_attribute": "new_value", "extra_attribute": "extra_value"})
		wg.Wait()
	}

	{
		for _, objectID := range objectIDs[:7] {
			res, err := index.DeleteObject(objectID)
			require.NoError(t, err)
			g.Collect(res)
		}
	}

	{
		res, err := index.DeleteObjects(objectIDs[7:])
		require.NoError(t, err)
		g.Collect(res)
	}

	require.NoError(t, g.Wait())

	{
		it, err := index.BrowseObjects()
		require.NoError(t, err)

		_, err = it.Next()
		require.Equal(t, io.EOF, err)
	}
}

func getObjectAndCompareWith(t *testing.T, wg *sync.WaitGroup, index *search.Index, expected map[string]string) {
	defer wg.Done()

	objectID, ok := expected["objectID"]
	require.True(t, ok)

	var found map[string]string
	err := index.GetObject(objectID, &found)
	require.NoError(t, err)
	require.Equal(t, expected, found)
}

func getObjectsAndCompareWith(t *testing.T, wg *sync.WaitGroup, index *search.Index, objectIDs []string, expected []map[string]string) {
	defer wg.Done()

	var objects []map[string]string
	err := index.GetObjects(objectIDs, &objects)
	require.NoError(t, err)
	require.ElementsMatch(t, expected, objects)
}
