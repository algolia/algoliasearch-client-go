package search_client

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/algolia/algoliasearch-client-go/cts"
	"github.com/stretchr/testify/require"
)

func TestCopyIndex(t *testing.T) {
	t.Parallel()
	client, index, indexName := cts.InitSearchClient1AndIndex(t)

	await := algolia.Await()

	{
		res, err := index.SaveObjects([]map[string]string{
			{"objectID": "one", "company": "apple"},
			{"objectID": "two", "company": "tesla"},
		})
		require.NoError(t, err)
		await.Collect(res)
	}

	expectedSettings := search.Settings{
		AttributesForFaceting: opt.AttributesForFaceting("company"),
	}

	{
		res, err := index.SetSettings(expectedSettings)
		require.NoError(t, err)
		await.Collect(res)
	}

	expectedSynonym := search.NewPlaceholder("google_placeholder", "<GOOG>", "Google", "GOOG")

	{
		res, err := index.SaveSynonym(expectedSynonym, true)
		require.NoError(t, err)
		await.Collect(res)
	}

	expectedRule := search.Rule{
		ObjectID:  "company_auto_faceting",
		Condition: search.RuleCondition{Anchoring: search.Contains, Pattern: "{facet:company}"},
		Consequence: search.RuleConsequence{
			Params: &search.RuleParams{
				AutomaticFacetFilters: []search.AutomaticFacetFilter{
					{Facet: "company"},
				},
			},
		},
	}

	{
		res, err := index.SaveRule(expectedRule, true)
		require.NoError(t, err)
		await.Collect(res)
	}

	require.NoError(t, await.Wait())

	{
		res, err := client.CopySettings(indexName, indexName+"_settings")
		require.NoError(t, err)
		await.Collect(res)
	}

	{
		res, err := client.CopyRules(indexName, indexName+"_rules")
		require.NoError(t, err)
		await.Collect(res)
	}

	{
		res, err := client.CopySynonyms(indexName, indexName+"_synonyms")
		require.NoError(t, err)
		await.Collect(res)
	}

	{
		res, err := client.CopyIndex(indexName, indexName+"_full_copy")
		require.NoError(t, err)
		await.Collect(res)
	}

	require.NoError(t, await.Wait())

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
			require.True(t, settings.Equal(expectedSettings))
		}

		if c.ShouldHaveSameRules {
			_, err := copiedIndex.GetRule(expectedRule.ObjectID)
			require.NoError(t, err)
		}

		if c.ShouldHaveSameSynonyms {
			_, err := copiedIndex.GetSynonym(expectedSynonym.ObjectID())
			require.NoError(t, err)
		}
	}
}
