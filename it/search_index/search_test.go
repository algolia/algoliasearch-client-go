package search_index

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	t.Parallel()
	_, index, _ := it.InitSearchClient1AndIndex(t)

	var taskIDs []int

	{
		res, err := index.AddObjects([]algoliasearch.Object{
			{"company": "Algolia", "name": "Julien Lemoine"},
			{"company": "Algolia", "name": "Nicolas Dessaigne"},
			{"company": "Amazon", "name": "Jeff Bezos"},
			{"company": "Apple", "name": "Steve Jobs"},
			{"company": "Apple", "name": "Steve Wozniak"},
			{"company": "Arista Networks", "name": "Jayshree Ullal"},
			{"company": "Google", "name": "Larry Page"},
			{"company": "Google", "name": "Rob Pike"},
			{"company": "Google", "name": "Sergueï Brin"},
			{"company": "Microsoft", "name": "Bill Gates"},
			{"company": "SpaceX", "name": "Elon Musk"},
			{"company": "Tesla", "name": "Elon Musk"},
			{"company": "Yahoo", "name": "Marissa Mayer"},
		})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index.SetSettings(algoliasearch.Map{
			"attributesForFaceting": []string{"searchable(company)"},
		})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	it.WaitTasks(t, index, taskIDs...)
	taskIDs = []int{}

	t.Run("Search", func(t *testing.T) {
		t.Parallel()
		res, err := index.Search("algolia", nil)
		require.NoError(t, err)
		require.Len(t, res.Hits, 2)
	})

	t.Run("SearchWithFacets", func(t *testing.T) {
		t.Parallel()
		{
			res, err := index.Search("elon", algoliasearch.Map{
				"facets":       "*",
				"facetFilters": "company:tesla",
			})
			require.NoError(t, err)
			require.Len(t, res.Hits, 1)
		}

		{
			res, err := index.Search("elon", algoliasearch.Map{
				"facets":  "*",
				"filters": "(company:tesla OR company:spacex)",
			})
			require.NoError(t, err)
			require.Len(t, res.Hits, 2)
		}
	})

	t.Run("SearchWithClickAnalytics", func(t *testing.T) {
		t.Parallel()
		res, err := index.Search("elon", algoliasearch.Map{
			"clickAnalytics": true,
		})
		require.NoError(t, err)
		require.NotEmpty(t, res.QueryID)
	})

	t.Run("SearchForFacetValues", func(t *testing.T) {
		t.Parallel()
		res, err := index.SearchForFacetValues("company", "a", nil)
		require.NoError(t, err)
		var foundFacets []string
		for _, f := range res.FacetHits {
			foundFacets = append(foundFacets, f.Value)
		}
		require.ElementsMatch(t, foundFacets, []string{
			"Algolia",
			"Amazon",
			"Apple",
			"Arista Networks",
		})
	})
}
