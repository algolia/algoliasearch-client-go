package query_suggestions

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/query_suggestions"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Parallel()
	searchClient, index1, indexName1 := cts.InitSearchClient1AndIndex(t)
	indexName2 := indexName1 + "_dev"
	index2 := searchClient.InitIndex(indexName2)
	querySuggestionsClient := cts.InitQuerySuggestionsClient1(t)

	// Create the two indices by adding a dummy object in each of them
	{
		g := wait.NewGroup()

		res, err := index1.SaveObject(map[string]string{"objectID": "one"})
		require.NoError(t, err)
		g.Collect(res)

		res, err = index2.SaveObject(map[string]string{"objectID": "one"})
		require.NoError(t, err)
		g.Collect(res)

		require.NoError(t, g.Wait())
	}

	indexName := "test_query_suggestion_index"

	indexConfig := query_suggestions.IndexConfiguration{
		IndexName: indexName,
		SourceIndices: []query_suggestions.SourceIndex{
			{
				IndexName:     "test_offerings_query_suggestions",
				AnalyticsTags: nil,
				Facets:        nil,
				MinHits:       func() *int { i := 4; return &i }(),
				MinLetters:    func() *int { i := 2; return &i }(),
				Generate:      nil,
			},
		},
		Languages: []string{"en"},
		Exclude:   nil,
	}

	t.Run("Create the query suggestion index", func(t *testing.T) {
		err := querySuggestionsClient.CreateConfig(indexConfig)
		require.NoError(t, err)
	})

	t.Run("Retrieve the query suggestion config", func(t *testing.T) {
		got, err := querySuggestionsClient.GetConfig(indexName)
		require.NoError(t, err)
		require.Equal(t, &indexConfig, got)
	})

	t.Run("Update the query suggestion index", func(t *testing.T) {
		indexConfig.Languages = []string{"ja"}
		err := querySuggestionsClient.UpdateConfig(indexConfig)
		require.NoError(t, err)

		got, err := querySuggestionsClient.GetConfig(indexName)
		require.NoError(t, err)
		require.Equal(t, &indexConfig, got)
	})

	t.Run("Retrieve all query suggestion configs", func(t *testing.T) {
		got, err := querySuggestionsClient.ListConfigs()
		require.NoError(t, err)
		require.Equal(t, &indexConfig, got[0])
	})

	t.Run("Delete the AB test", func(t *testing.T) {
		err := querySuggestionsClient.DeleteConfig(indexName)
		require.NoError(t, err)

		_, err = querySuggestionsClient.GetConfig(indexName)
		require.Error(t, err)
	})
}
