package insights_client

import (
	"net/http"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/debug"
	"github.com/algolia/algoliasearch-client-go/algolia/insights"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/cts"
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

	debug.Enable()

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
		res, err := insightsUserClient.ConvertedObjectIDsAfterSearch("foo", indexName, []string{"one", "two"}, queryID)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Status)
	}

	{
		res, err := insightsUserClient.ConvertedFilters("foo", indexName, []string{"filter:foo", "filter:bar"})
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
