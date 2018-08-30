package search_index

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
	"github.com/stretchr/testify/require"
)

func TestSettings(t *testing.T) {
	t.Parallel()
	_, index, indexName := it.InitSearchClient1AndIndex(t)

	// Add dummy object
	{
		res, err := index.AddObject(algoliasearch.Object{"objectID": "one", "attribute": "value"})
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)
	}

	expected := algoliasearch.Map{
		// Attributes
		"searchableAttributes":    []string{"attribute1", "attribute2", "attribute3", "ordered(attribute4)", "unordered(attribute5)"},
		"attributesForFaceting":   []string{"attribute1", "filterOnly(attribute2)", "searchable(attribute3)"},
		"unretrievableAttributes": []string{"attribute1", "attribute2"},
		"attributesToRetrieve":    []string{"attribute3", "attribute4"},
		// Ranking
		"ranking":       []string{"asc(attribute1)", "desc(attribute2)", "attribute", "custom", "exact", "filters", "geo", "proximity", "typo", "words"},
		"customRanking": []string{"asc(attribute1)", "desc(attribute2)"},
		"replicas":      []string{indexName + "_replica1", indexName + "_replica2"},
		// Faceting
		"maxValuesPerFacet": 100,
		"sortFacetValuesBy": "count",
		// Highlighting / Snippeting
		"attributesToHighlight":             []string{"attribute1", "attribute2"},
		"attributesToSnippet":               []string{"attribute1:10", "attribute2:8"},
		"highlightPreTag":                   "<strong>",
		"highlightPostTag":                  "</strong>",
		"snippetEllipsisText":               " and so on.",
		"restrictHighlightAndSnippetArrays": true,
		// Pagination
		"hitsPerPage":         42,
		"paginationLimitedTo": 43,
		// Typos
		"minWordSizefor1Typo":              2,
		"minWordSizefor2Typos":             6,
		"typoTolerance":                    false,
		"allowTyposOnNumericTokens":        false,
		"ignorePlurals":                    true,
		"disableTypoToleranceOnAttributes": []string{"attribute1", "attribute2"},
		"disableTypoToleranceOnWords":      []string{"word1", "word2"},
		"separatorsToIndex":                "()[]",
		// Query strategy
		"queryType":                 "prefixNone",
		"removeWordsIfNoResults":    "allOptional",
		"advancedSyntax":            true,
		"optionalWords":             []string{"word1", "word2"},
		"removeStopWords":           true,
		"disablePrefixOnAttributes": []string{"attribute1", "attribute2"},
		"disableExactOnAttributes":  []string{"attribute1", "attribute2"},
		"exactOnSingleWordQuery":    "word",
		// Query rules
		"enableRules": false,
		// Performance
		"numericAttributesForFiltering":  []string{"attribute1", "attribute2"},
		"allowCompressionOfIntegerArray": true,
		// Advanced
		"attributeForDistinct":       "attribute1",
		"distinct":                   2,
		"replaceSynonymsInHighlight": false,
		"minProximity":               7,
		"responseFields":             []string{"hits", "hitsPerPage"},
		"maxFacetHits":               100,
		"camelCaseAttributes":        []string{"attribute1", "attribute2"},
		"decompoundedAttributes":     map[string][]string{"de": []string{"attribute1", "attribute2"}, "fi": []string{"attribute3"}},
		"keepDiacriticsOnCharacters": "øé",
	}

	{
		res, err := index.SetSettings(expected)
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)

		settings, err := index.GetSettings()
		require.NoError(t, err)
		settingsAsMap := settings.ToMap()
		require.Equal(t, expected, settingsAsMap)

		res, err = index.SetSettings(settingsAsMap)
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)

		settings, err = index.GetSettings()
		require.NoError(t, err)
		settingsAsMap = settings.ToMap()
		require.Equal(t, expected, settingsAsMap)
	}

	expected["typoTolerance"] = "min"
	expected["ignorePlurals"] = []string{"en", "fr"}
	expected["removeStopWords"] = []string{"en", "fr"}
	expected["distinct"] = true

	{
		res, err := index.SetSettings(expected)
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)

		settings, err := index.GetSettings()
		require.NoError(t, err)
		settingsAsMap := settings.ToMap()
		require.Equal(t, expected, settingsAsMap)

		res, err = index.SetSettings(settingsAsMap)
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)

		settings, err = index.GetSettings()
		require.NoError(t, err)
		settingsAsMap = settings.ToMap()
		require.Equal(t, expected, settingsAsMap)
	}
}
