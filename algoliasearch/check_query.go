package algoliasearch

func checkQuery(query map[string]interface{}, ignore ...string) error {
Outer:
	for k, v := range query {
		// Continue if `k` is to be ignored.
		for _, s := range ignore {
			if s == k {
				continue Outer
			}
		}

		switch v.(type) {
		case string:
			if k != "query" && k != "queryType" &&
				k != "typoTolerance" &&
				k != "removeWordsIfNoResults" &&
				k != "restrictSearchableAttributes" &&
				k != "highlightPreTag" &&
				k != "highlightPostTag" &&
				k != "snippetEllipsisText" &&
				k != "filters" &&
				k != "analyticsTags" &&
				k != "optionalWords" &&
				k != "disableTypoToleranceOnAttributes" &&
				k != "attributesToRetrieve" &&
				k != "attributesToHighlight" &&
				k != "numericFilters" &&
				k != "tagFilters" &&
				k != "facets" &&
				k != "facetFilters" &&
				k != "attributesToSnippet" &&
				k != "aroundLatLng" &&
				k != "insideBoundingBox" &&
				k != "insidePolygon" {
				return invalidParameter(k)
			}

		case int64:
			if k != "minWordSizefor1Typo" &&
				k != "minWordSizefor2Typos" &&
				k != "minProximity" &&
				k != "page" &&
				k != "hitsPerPage" &&
				k != "getRankingInfo" &&
				k != "distinct" &&
				k != "maxValuesPerFacet" &&
				k != "aroundRadius" &&
				k != "aroundPrecision" &&
				k != "minimumAroundRadius" {
				return invalidParameter(k)
			}

		case bool:
			if k != "allowTyposOnNumericTokens" &&
				k != "ignorePlurals" &&
				k != "advancedSyntax" &&
				k != "analytics" &&
				k != "synonyms" &&
				k != "replaceSynonymsInHighlight" &&
				k != "removeStopWords" &&
				k != "aroundLatLngViaIP" {
				return invalidParameter(k)
			}

		default:
			return invalidParameter(k)
		}
	}

	return nil
}
