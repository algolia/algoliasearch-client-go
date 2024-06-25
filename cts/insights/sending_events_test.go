package insights

import (
	"net/http"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/insights"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestSendingEvents(t *testing.T) {
	t.Parallel()
	_, index, indexName := cts.InitSearchClient1AndIndex(t)

	{
		res, err := index.SaveObjects([]map[string]string{
			{"objectID": "one"},
			{"objectID": "two"},
		})
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	insightsClient := cts.InitInsightsClient(t)

	{
		res, err := insightsClient.SendEvent(insights.Event{
			EventType: insights.EventTypeClick,
			EventName: "foo",
			Index:     indexName,
			UserToken: "bar",
			ObjectIDs: []string{"one", "two"},
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsClient.SendEvents([]insights.Event{
			{
				EventType: insights.EventTypeClick,
				EventName: "foo",
				Index:     indexName,
				UserToken: "bar",
				ObjectIDs: []string{"one", "two"},
			},
			{
				EventType: insights.EventTypeClick,
				EventName: "foo",
				Index:     indexName,
				UserToken: "bar",
				ObjectIDs: []string{"one", "two"},
			},
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	insightsUserClient := insightsClient.User("bar")

	{
		res, err := insightsUserClient.ClickedObjectIDs("foo", indexName, []string{"one", "two"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	var queryID string

	{
		res, err := index.Search("", opt.ClickAnalytics(true))
		require.NoError(t, err)
		queryID = res.QueryID
	}

	{
		res, err := insightsUserClient.ClickedObjectIDsAfterSearch("foo", indexName, []string{"one", "two"}, []int{1, 2}, queryID)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ClickedFilters("foo", indexName, []string{"filter:foo", "filter:bar"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ConvertedObjectIDs("foo", indexName, []string{"one", "two"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ConvertedObjectIDsPurchase("foo", indexName, []string{"one", "two"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ConvertedObjectIDsAddToCart("foo", indexName, []string{"one", "two"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ConvertedObjectIDsAfterSearch("foo", indexName, []string{"one", "two"}, queryID)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ConvertedObjectIDsAfterSearchPurchase("foo", indexName, []string{"one", "two"}, queryID)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ConvertedObjectIDsAfterSearchAddToCart("foo", indexName, []string{"one", "two"}, queryID)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ConvertedFilters("foo", indexName, []string{"filter:foo", "filter:bar"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ConvertedFiltersPurchase("foo", indexName, []string{"filter:foo", "filter:bar"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ConvertedFiltersAddToCart("foo", indexName, []string{"filter:foo", "filter:bar"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ViewedObjectIDs("foo", indexName, []string{"one", "two"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ViewedFilters("foo", indexName, []string{"filter:foo", "filter:bar"})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}
}
