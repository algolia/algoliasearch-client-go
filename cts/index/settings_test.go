package index

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/search"
	"github.com/algolia/algoliasearch-client-go/cts"
	"github.com/stretchr/testify/require"
)

func TestSettings(t *testing.T) {
	t.Parallel()
	_, index, indexName := cts.InitSearchClient1AndIndex(t)

	await := algolia.Await()

	{
		res, err := index.SaveObject(map[string]string{"objectID": "one", "attribute": "value"})
		require.NoError(t, err)
		await.Collect(res)
	}

	expected := search.Settings{
		SearchableAttributes:              opt.SearchableAttributes("attribute1", "attribute2", "attribute3", "ordered(attribute4)", "unordered(attribute5)"),
		AttributesForFaceting:             opt.AttributesForFaceting("attribute1", "filterOnly(attribute2)", "searchable(attribute3)"),
		UnretrievableAttributes:           opt.UnretrievableAttributes("attribute1", "attribute2"),
		AttributesToRetrieve:              opt.AttributesToRetrieve("attribute3", "attribute4"),
		Ranking:                           opt.Ranking("asc(attribute1)", "desc(attribute2)", "attribute", "custom", "exact", "filters", "geo", "proximity", "typo", "words"),
		CustomRanking:                     opt.CustomRanking("asc(attribute1)", "desc(attribute2)"),
		Replicas:                          opt.Replicas(indexName+"_replica1", indexName+"_replica2"),
		MaxValuesPerFacet:                 opt.MaxValuesPerFacet(100),
		SortFacetValuesBy:                 opt.SortFacetValuesBy("count"),
		AttributesToHighlight:             opt.AttributesToHighlight("attribute1", "attribute2"),
		AttributesToSnippet:               opt.AttributesToSnippet("attribute1:10", "attribute2:8"),
		HighlightPreTag:                   opt.HighlightPreTag("<strong>"),
		HighlightPostTag:                  opt.HighlightPostTag("</strong>"),
		SnippetEllipsisText:               opt.SnippetEllipsisText(" and so on."),
		RestrictHighlightAndSnippetArrays: opt.RestrictHighlightAndSnippetArrays(true),
		HitsPerPage:                       opt.HitsPerPage(42),
		PaginationLimitedTo:               opt.PaginationLimitedTo(43),
		MinWordSizefor1Typo:               opt.MinWordSizefor1Typo(2),
		MinWordSizefor2Typos:              opt.MinWordSizefor2Typos(6),
		TypoTolerance:                     opt.TypoTolerance(false),
		AllowTyposOnNumericTokens:         opt.AllowTyposOnNumericTokens(false),
		DisableTypoToleranceOnAttributes:  opt.DisableTypoToleranceOnAttributes("attribute1", "attribute2"),
		DisableTypoToleranceOnWords:       opt.DisableTypoToleranceOnWords("word1", "word2"),
		SeparatorsToIndex:                 opt.SeparatorsToIndex("()[]"),
		IgnorePlurals:                     opt.IgnorePlurals(true),
		RemoveStopWords:                   opt.RemoveStopWords(true),
		CamelCaseAttributes:               opt.CamelCaseAttributes("attribute1", "attribute2"),
		DecompoundedAttributes:            opt.DecompoundedAttributes(map[string][]string{"de": {"attribute1", "attribute2"}, "fi": {"attribute3"}}),
		KeepDiacriticsOnCharacters:        opt.KeepDiacriticsOnCharacters("øé"),
		QueryLanguages:                    opt.QueryLanguages("fr", "en"),
		QueryType:                         opt.QueryType("prefixNone"),
		RemoveWordsIfNoResults:            opt.RemoveWordsIfNoResults("allOptional"),
		AdvancedSyntax:                    opt.AdvancedSyntax(true),
		OptionalWords:                     opt.OptionalWords("word1", "word2"),
		DisablePrefixOnAttributes:         opt.DisablePrefixOnAttributes("attribute1", "attribute2"),
		DisableExactOnAttributes:          opt.DisableExactOnAttributes("attribute1", "attribute2"),
		ExactOnSingleWordQuery:            opt.ExactOnSingleWordQuery("word"),
		AlternativesAsExact:               opt.AlternativesAsExact("ignorePlurals"),
		AdvancedSyntaxFeatures:            opt.AdvancedSyntaxFeatures("exactPhrase"),
		EnableRules:                       opt.EnableRules(false),
		NumericAttributesForFiltering:     opt.NumericAttributesForFiltering("attribute1", "attribute2"),
		AllowCompressionOfIntegerArray:    opt.AllowCompressionOfIntegerArray(true),
		AttributeForDistinct:              opt.AttributeForDistinct("attribute1"),
		Distinct:                          opt.DistinctOf(2),
		ReplaceSynonymsInHighlight:        opt.ReplaceSynonymsInHighlight(false),
		MinProximity:                      opt.MinProximity(7),
		ResponseFields:                    opt.ResponseFields("hits", "hitsPerPage"),
		MaxFacetHits:                      opt.MaxFacetHits(100),
	}

	{
		res, err := index.SetSettings(expected)
		require.NoError(t, err)
		await.Collect(res)
	}

	require.NoError(t, await.Wait())

	{
		settings, err := index.GetSettings()
		require.NoError(t, err)
		require.True(t, settings.Equal(expected))

		res, err := index.SetSettings(settings)
		require.NoError(t, err)
		require.NoError(t, res.Wait())

		settings, err = index.GetSettings()
		require.NoError(t, err)
		require.True(t, settings.Equal(expected))
	}

	expected.TypoTolerance = opt.TypoToleranceMin()
	expected.IgnorePlurals = opt.IgnorePluralsFor("en", "fr")
	expected.RemoveStopWords = opt.RemoveStopWordsFor("en", "fr")
	expected.Distinct = opt.Distinct(true)

	{
		res, err := index.SetSettings(expected)
		require.NoError(t, err)
		require.NoError(t, res.Wait())

		settings, err := index.GetSettings()
		require.NoError(t, err)
		require.True(t, settings.Equal(expected))

		res, err = index.SetSettings(settings)
		require.NoError(t, err)
		require.NoError(t, res.Wait())

		settings, err = index.GetSettings()
		require.NoError(t, err)
		require.True(t, settings.Equal(expected))
	}
}
