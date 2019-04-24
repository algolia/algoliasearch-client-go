package index

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/algolia/algoliasearch-client-go/cts"
	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	t.Parallel()
	_, index, _ := cts.InitSearchClient1AndIndex(t)

	await := algolia.Await()

	{
		res, err := index.SaveObjects([]map[string]string{
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
		}, opt.AutoGenerateObjectIDIfNotExist(true))
		require.NoError(t, err)
		await.Collect(res)
	}

	{
		res, err := index.SetSettings(search.Settings{
			AttributesForFaceting: opt.AttributesForFaceting("searchable(company)"),
		})
		require.NoError(t, err)
		await.Collect(res)
	}

	require.NoError(t, await.Wait())

	{
		res, err := index.Search("algolia", nil)
		require.NoError(t, err)
		require.Len(t, res.Hits, 2)
	}

	{
		res, err := index.Search("elon",
			opt.Facets("*"),
			opt.FacetFilter("company:tesla"),
		)
		require.NoError(t, err)
		require.Len(t, res.Hits, 1)
	}

	{
		res, err := index.Search("elon",
			opt.Facets("*"),
			opt.Filters("(company:tesla OR company:spacex)"),
		)
		require.NoError(t, err)
		require.Len(t, res.Hits, 2)
	}

	{
		res, err := index.Search("elon",
			opt.ClickAnalytics(true),
		)
		require.NoError(t, err)
		require.NotEmpty(t, res.QueryID)
	}

	{
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
	}
}
