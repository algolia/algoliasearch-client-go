package search_client

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
)

func TestCopyIndex(t *testing.T) {
	t.Parallel()
	client, index, indexName := it.InitSearchClient1AndIndex(t)

	var taskIDs []int

	{
		res, err := index.AddObjects([]algoliasearch.Object{
			{"objectID": "one", "company": "apple"},
			{"objectID": "two", "company": "tesla"},
		})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	expectedSettings := algoliasearch.Map{
		"attributesForFaceting": []string{"company"},
	}

	{
		res, err := index.SetSettings(expectedSettings)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	expectedSynonym := algoliasearch.NewPlaceholderSynonym(
		"google_placeholder",
		"<GOOG>",
		[]string{"Google", "GOOG"},
	)

	{
		res, err := index.SaveSynonym(expectedSynonym, true)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	expectedRule := algoliasearch.Rule{
		ObjectID: "company_auto_faceting",
		Condition: algoliasearch.NewSimpleRuleCondition(
			algoliasearch.Contains,
			"{facet:company}",
		),
		Consequence: algoliasearch.RuleConsequence{
			Params: algoliasearch.Map{
				"automaticFacetFilters": []string{"company"},
			},
		},
	}

	{
		res, err := index.SaveRule(expectedRule, true)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	it.WaitTasks(t, index, taskIDs...)
	taskIDs = []int{}

	{
		res, err := client.CopySettings(indexName, indexName+"_settings")
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := client.CopyRules(indexName, indexName+"_rules")
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := client.CopySynonyms(indexName, indexName+"_synonyms")
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := client.CopyIndex(indexName, indexName+"_full_copy")
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	it.WaitTasks(t, index, taskIDs...)
	taskIDs = []int{}

	for _, c := range []struct {
		IndexName              string
		ShouldHaveSameSettings bool
		ShouldHaveSameRules    bool
		ShouldHaveSameSynonyms bool
	}{
		{IndexName: indexName + "_settings", ShouldHaveSameSettings: true},
		{IndexName: indexName + "_rules", ShouldHaveSameRules: true},
		{IndexName: indexName + "_synonyms", ShouldHaveSameSynonyms: true},
		{IndexName: indexName + "_full_copy", ShouldHaveSameSettings: true, ShouldHaveSameRules: true, ShouldHaveSameSynonyms: true},
	} {
		copiedIndex := client.InitIndex(c.IndexName)

		if c.ShouldHaveSameSettings {
			settings, err := copiedIndex.GetSettings()
			require.NoError(t, err)
			require.Equal(t, expectedSettings["attributesForFaceting"], settings.AttributesForFaceting)
		}

		if c.ShouldHaveSameRules {
			_, err := copiedIndex.GetRule(expectedRule.ObjectID)
			require.NoError(t, err)
		}

		if c.ShouldHaveSameSynonyms {
			_, err := copiedIndex.GetSynonym(expectedSynonym.ObjectID)
			require.NoError(t, err)
		}
	}
}
