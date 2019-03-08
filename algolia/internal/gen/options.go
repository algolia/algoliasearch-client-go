//+build ignore

package main

type Kind int

const (
	Settings Kind = 1
	Search   Kind = 2
	Other    Kind = 4
)

func isSettings(k Kind) bool { return k&Settings != 0 }
func isSearch(k Kind) bool   { return k&Search != 0 }

type Option struct {
	Name                          string
	Kind                          Kind
	DefaultValue                  interface{} // Leave nil to prevent generation (used for hand-written options)
	BackwardCompatibleAlternative string
}

// Full reference of all settings and search parameters can be found here:
// https://www.algolia.com/doc/api-reference/api-parameters/

var options = []Option{
	// Attributes
	{"searchableAttributes", Settings | Search, []string{}, "attributesToIndex"},
	{"attributesForFaceting", Settings | Search, []string{}, ""},
	{"unretrievableAttributes", Settings | Search, []string{}, ""},
	{"attributesToRetrieve", Settings | Search, []string{"*"}, ""},
	{"restrictSearchableAttributes", Search, []string{}, ""},

	// Ranking
	{"ranking", Settings | Search, []string{"typo", "geo", "words", "filters", "proximity", "attribute", "exact", "custom"}, ""},
	{"customRanking", Settings | Search, []string{}, ""},
	{"replicas", Settings | Search, []string{}, "slaves"},

	// Filtering
	{"filters", Search, "attribute", ""},
	{"facetFilters", Search, nil, ""},
	{"optionalFilters", Search, nil, ""},
	{"numericFilters", Search, nil, ""},
	{"tagFilters", Search, nil, ""},
	{"sumOrFiltersScores", Search, false, ""},

	// Faceting
	{"facets", Search, []string{}, ""},
	{"maxValuesPerFacet", Settings | Search, 100, ""},
	{"facetingAfterDistinct", Search, false, ""},
	{"sortFacetValuesBy", Settings | Search, "count", ""},

	// Highlighting / Snippeting
	{"attributesToHighlight", Search | Settings, []string{}, ""},
	{"attributesToSnippet", Search | Settings, []string{}, ""},
	{"highlightPreTag", Search | Settings, "<em>", ""},
	{"highlightPostTag", Search | Settings, "</em>", ""},
	{"snippetEllipsisText", Search | Settings, "…", ""},
	{"restrictHighlightAndSnippetArrays", Settings | Search, false, ""},

	// Pagination
	{"page", Search, 0, ""},
	{"hitsPerPage", Settings | Search, 20, ""},
	{"offset", Search, 0, ""},
	{"length", Search, 0, ""},
	{"paginationLimitedTo", Settings, 1000, ""},

	// Typos
	{"minWordSizefor1Typo", Settings | Search, 4, ""},
	{"minWordSizefor2Typos", Settings | Search, 8, ""},
	{"typoTolerance", Settings | Search, nil, ""},
	{"allowTyposOnNumericTokens", Settings | Search, true, ""},
	{"disableTypoToleranceOnAttributes", Settings | Search, []string{}, ""},
	{"disableTypoToleranceOnWords", Settings, []string{}, ""},
	{"separatorsToIndex", Settings, "", ""},

	// Geo Search
	{"aroundLatLng", Search, "", ""},
	{"aroundLatLngViaIP", Search, false, ""},
	{"aroundRadius", Search, nil, ""},
	{"aroundPrecision", Search, 1, ""},
	{"minimumAroundRadius", Search, 0, ""},
	{"insideBoundingBox", Search, nil, ""},
	{"insidePolygon", Search, nil, ""},

	// Languages
	{"ignorePlurals", Settings | Search, nil, ""},
	{"removeStopWords", Settings | Search, nil, ""},
	{"camelCaseAttributes", Settings, []string{}, ""},
	{"decompoundedAttributes", Settings, map[string][]string{}, ""},
	{"keepDiacriticsOnCharacters", Settings, "", ""},
	{"queryLanguages", Settings | Search, []string{}, ""},

	// Query strategy
	{"queryType", Settings | Search, "prefixLast", ""},
	{"removeWordsIfNoResults", Settings | Search, "none", ""},
	{"advancedSyntax", Settings | Search, false, ""},
	{"optionalWords", Settings | Search, []string{}, ""},
	{"disablePrefixOnAttributes", Settings, []string{}, ""},
	{"disableExactOnAttributes", Settings | Search, []string{}, ""},
	{"exactOnSingleWordQuery", Settings | Search, "attribute", ""},
	{"alternativesAsExact", Settings | Search, []string{"ignorePlurals", "singleWordSynonym"}, ""},
	{"advancedSyntaxFeatures", Settings | Search, []string{"exactPhrase", "excludeWords"}, ""},

	// Query rules
	{"enableRules", Settings | Search, true, ""},
	{"ruleContexts", Search, []string{}, ""},

	// Personalization
	{"enablePersonalization", Search, false, ""},

	// Performance
	{"numericAttributesForFiltering", Settings, []string{}, "numericAttributesToIndex"},
	{"allowCompressionOfIntegerArray", Settings, false, ""},

	// Advanced
	{"attributeForDistinct", Settings, "", ""},
	{"distinct", Settings | Search, 0, ""},
	{"getRankingInfo", Search, false, ""},
	{"clickAnalytics", Search, false, ""},
	{"analytics", Search, true, ""},
	{"analyticsTags", Search, []string{}, ""},
	{"synonyms", Search, true, ""},
	{"replaceSynonymsInHighlight", Settings | Search, true, ""},
	{"minProximity", Settings | Search, 1, ""},
	{"responseFields", Settings | Search, []string{"*"}, ""},
	{"maxFacetHits", Settings | Search, 10, ""},
	{"percentileComputation", Search, true, ""},

	// Other
	{"query", Other, "", ""},
	{"autoGenerateObjectIDIfNotExist", Other, false, ""},
	// TODO: fix inconsistency between `clearExistingRules` / `replaceExistingSynonyms`
	{"clearExistingRules", Other, false, ""},
	{"replaceExistingSynonyms", Other, false, ""},
	{"type", Other, "", ""},
	{"createIfNotExists", Other, false, ""},
	{"forwardToReplicas", Other, false, ""},
	{"anchoring", Other, "", ""},
	{"extraHeaders", Other, map[string]string{}, ""},
	{"extraURLParams", Other, map[string]string{}, ""},
}
