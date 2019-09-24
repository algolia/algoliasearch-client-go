package index

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
)

func TestSearch(t *testing.T) {
	t.Parallel()
	_, index, _ := cts.InitSearchClient1AndIndex(t)

	g := wait.NewGroup()

	{
		res, err := index.SaveObjects([]map[string]string{
			{"company": "Algolia", "name": "Julien Lemoine", "objectID": "julien-lemoine"},
			{"company": "Algolia", "name": "Nicolas Dessaigne", "objectID": "nicolas-dessaigne"},
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
		g.Collect(res)
	}

	{
		res, err := index.SetSettings(search.Settings{
			AttributesForFaceting: opt.AttributesForFaceting("searchable(company)"),
		})
		require.NoError(t, err)
		g.Collect(res)
	}

	require.NoError(t, g.Wait())

	{
		res, err := index.Search("algolia", nil)
		require.NoError(t, err)
		require.Len(t, res.Hits, 2)
		require.Equal(t, 0, res.GetObjectPosition("nicolas-dessaigne"))
		require.Equal(t, 1, res.GetObjectPosition("julien-lemoine"))
		require.Equal(t, -1, res.GetObjectPosition(""))
	}

	{
		filterFunc := func(object map[string]interface{}) bool { return false }
		obj, err := index.FindObject(filterFunc, "", true)
		require.Error(t, err, "object %#v was found but it should not", obj)

		filterFunc = func(object map[string]interface{}) bool { return true }
		obj, err = index.FindObject(filterFunc, "", true)
		require.NoError(t, err)
		require.Equal(t, 0, obj.Position)
		require.Equal(t, 0, obj.Page)

		filterFunc = func(object map[string]interface{}) bool {
			itf, ok := object["company"]
			if !ok {
				return false
			}
			company, ok := itf.(string)
			return ok && company == "Apple"
		}

		obj, err = index.FindObject(filterFunc, "algolia", true)
		require.Error(t, err, "object %#v was found but it should not", obj)

		obj, err = index.FindObject(filterFunc, "", false, opt.HitsPerPage(5))
		require.Error(t, err, "object %#v was found but it should not", obj)

		obj, err = index.FindObject(filterFunc, "", true, opt.HitsPerPage(5))
		require.NoError(t, err)
		require.Equal(t, 0, obj.Position)
		require.Equal(t, 2, obj.Page)
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
